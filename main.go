package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/rest"
	"log"
	"mini-k8s-proxy/pkg/apis/miniproxy/v1alpha1"
	clientset "mini-k8s-proxy/pkg/generated/clientset/versioned"
	"mini-k8s-proxy/pkg/generated/informers/externalversions"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type ProxyRouteSpec struct {
	V map[string]v1alpha1.ProxyRouteSpec
	sync.RWMutex
}

var prs = &ProxyRouteSpec{
	V: make(map[string]v1alpha1.ProxyRouteSpec, 0),
}

type resourceEventHandler struct {
	Ev chan<- interface{}
}

func (reh *resourceEventHandler) OnAdd(obj interface{}) {
	eventHandlerFunc(reh.Ev, obj)
}

func (reh *resourceEventHandler) OnUpdate(oldObj, newObj interface{}) {
	eventHandlerFunc(reh.Ev, newObj)
}

func (reh *resourceEventHandler) OnDelete(obj interface{}) {
	eventHandlerFunc(reh.Ev, obj)
}

func eventHandlerFunc(events chan<- interface{}, obj interface{}) {
	select {
	case events <- obj:
	default:
	}
}

func main() {
	ctx := context.Background()

	eventCh := make(chan interface{}, 1)
	// 采用缓冲大小为 1 的通道方式来处理 CRD 事件
	eventHandler := &resourceEventHandler{Ev: eventCh}

	// 作为测试，可以直接使用 kubeconfig 连接 k8s，实际部署使用 InClusterConfig 模式
	//cfg, err := clientcmd.BuildConfigFromFlags("", "tmp/config")
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	client, err := clientset.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	// 构建 k8s Crd Informer 实例
	factoryCrd := externalversions.NewSharedInformerFactoryWithOptions(
		client,
		10*time.Minute,
	)
	// 注册 Informer 事件处理
	factoryCrd.Miniproxy().V1alpha1().ProxyRoutes().Informer().AddEventHandler(eventHandler)
	// 启动 Informer
	factoryCrd.Start(ctx.Done())
	// 等待首次缓存同步
	for t, ok := range factoryCrd.WaitForCacheSync(ctx.Done()) {
		if !ok {
			panic(fmt.Errorf("timed out waiting for controller caches to sync %s", t.String()))
		}
	}
	go startServer()

	for {
		select {
		case _, ok := <-eventCh:
			if !ok {
				continue
			}
			// 从 Lister 缓存获取 CRD 资源对象
			proxyRoutes, err := factoryCrd.Miniproxy().V1alpha1().ProxyRoutes().Lister().List(labels.Everything())
			if err != nil {
				log.Println(err.Error())
				continue
			}
			// 清空本地缓存并重新放入
			prs.Lock()
			prs.V = make(map[string]v1alpha1.ProxyRouteSpec, 0)
			for _, proxyRoute := range proxyRoutes {
				fmt.Printf("%+v\n", proxyRoute)
				prs.V[proxyRoute.Spec.Host] = proxyRoute.Spec
			}
			prs.Unlock()
		}
	}
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Any("/*any", handler)
	log.Fatalln(r.Run(":80"))
}

func handler(c *gin.Context) {
	prs.RLock()
	defer prs.RUnlock()
	if proxyRouteSpec, ok := prs.V[c.Request.Host]; ok {
		u := ""
		if proxyRouteSpec.Scheme {
			u += "https://"
		} else {
			u += "http://"
		}
		if proxyRouteSpec.Namespace != "" {
			u += proxyRouteSpec.ServiceName + "." + proxyRouteSpec.Namespace
		} else {
			u += proxyRouteSpec.ServiceName
		}
		if proxyRouteSpec.Port != 0 {
			u += fmt.Sprintf(":%d", proxyRouteSpec.Port)
		}
		log.Println("代理地址: ", u)
		proxyUrl, err := url.Parse(u)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		proxyServer(c, proxyUrl)
	} else {
		c.String(http.StatusNotFound, "404")
	}
}

// 代理转发
func proxyServer(c *gin.Context, proxyUrl *url.URL) {
	proxy := &httputil.ReverseProxy{
		Director: func(outReq *http.Request) {
			u := outReq.URL
			outReq.URL = proxyUrl
			if outReq.RequestURI != "" {
				parsedURL, err := url.ParseRequestURI(outReq.RequestURI)
				if err == nil {
					u = parsedURL
				}
			}

			outReq.URL.Path = u.Path
			outReq.URL.RawPath = u.RawPath
			outReq.URL.RawQuery = u.RawQuery
			outReq.RequestURI = "" // Outgoing request should not have RequestURI

			outReq.Proto = "HTTP/1.1"
			outReq.ProtoMajor = 1
			outReq.ProtoMinor = 1

			if _, ok := outReq.Header["User-Agent"]; !ok {
				outReq.Header.Set("User-Agent", "")
			}

			// Even if the websocket RFC says that headers should be case-insensitive,
			// some servers need Sec-WebSocket-Key, Sec-WebSocket-Extensions, Sec-WebSocket-Accept,
			// Sec-WebSocket-Protocol and Sec-WebSocket-Version to be case-sensitive.
			// https://tools.ietf.org/html/rfc6455#page-20
			outReq.Header["Sec-WebSocket-Key"] = outReq.Header["Sec-Websocket-Key"]
			outReq.Header["Sec-WebSocket-Extensions"] = outReq.Header["Sec-Websocket-Extensions"]
			outReq.Header["Sec-WebSocket-Accept"] = outReq.Header["Sec-Websocket-Accept"]
			outReq.Header["Sec-WebSocket-Protocol"] = outReq.Header["Sec-Websocket-Protocol"]
			outReq.Header["Sec-WebSocket-Version"] = outReq.Header["Sec-Websocket-Version"]
			delete(outReq.Header, "Sec-Websocket-Key")
			delete(outReq.Header, "Sec-Websocket-Extensions")
			delete(outReq.Header, "Sec-Websocket-Accept")
			delete(outReq.Header, "Sec-Websocket-Protocol")
			delete(outReq.Header, "Sec-Websocket-Version")
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		ErrorHandler: func(w http.ResponseWriter, request *http.Request, err error) {
			statusCode := http.StatusInternalServerError
			w.WriteHeader(statusCode)
			w.Write([]byte(http.StatusText(statusCode)))
		},
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

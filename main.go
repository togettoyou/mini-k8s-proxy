package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	clientset "mini-k8s-proxy/pkg/generated/clientset/versioned"
	"mini-k8s-proxy/pkg/generated/informers/externalversions"
	"time"
)

const resyncPeriod = 10 * time.Minute

func main() {
	ctx := context.Background()
	eventCh := make(chan interface{}, 1)
	eventHandler := &resourceEventHandler{Ev: eventCh}

	// 使用 kubeconfig 连接 k8s client
	cfg, err := clientcmd.BuildConfigFromFlags("", "tmp/config")
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
		resyncPeriod,
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
				return
			}
			proxyRoutes, err := factoryCrd.Miniproxy().V1alpha1().ProxyRoutes().Lister().List(labels.Everything())
			if err != nil {
				log.Println(err.Error())
				continue
			}
			for _, proxyRoute := range proxyRoutes {
				fmt.Printf("%+v\n", proxyRoute)
			}

		}
	}
}

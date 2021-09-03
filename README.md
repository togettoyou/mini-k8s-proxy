# mini-k8s-proxy

# 部署

```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: proxyroutes.miniproxy.togettoyou.com
spec:
  group: miniproxy.togettoyou.com
  names:
    kind: ProxyRoute
    shortNames:
      - pr
    plural: proxyroutes
    singular: proxyroute
  scope: Namespaced
  versions:
    - name: v1alpha1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                host:
                  type: string
                serviceName:
                  type: string
                namespace:
                  type: string
                port:
                  type: integer
                scheme:
                  type: boolean
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mini-k8s-proxy
spec:
  selector:
    matchLabels:
      app: mini-k8s-proxy
  replicas: 1
  template:
    metadata:
      labels:
        app: mini-k8s-proxy
    spec:
      containers:
        - name: mini-k8s-proxy
          image: togettoyou/mini-k8s-proxy:latest
          ports:
            - containerPort: 80

---
apiVersion: v1
kind: Service
metadata:
  name: mini-k8s-proxy-service
spec:
  ports:
    - port: 80
      targetPort: 80
  selector:
    app: mini-k8s-proxy
  type: NodePort
```

```shell
kubectl apply -f mini-k8s-proxy.yaml
```

# 使用

```yaml
apiVersion: miniproxy.togettoyou.com/v1alpha1
kind: ProxyRoute
metadata:
  name: example-proxyroute
spec:
  # 监听域名
  host: whoami.togettoyou.com
  # 假设你有一个 whomai 的 service，位于 default 命名空间，容器内部端口为 80 ，http 协议
  serviceName: whoami
  namespace: default
  port: 80
  scheme: false
```

```shell
kubectl apply -f example.yaml
```
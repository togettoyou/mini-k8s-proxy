# mini-k8s-proxy

# 开发教程

[利用 CRD 实现一个 mini-k8s-proxy](https://mp.weixin.qq.com/s/SXF8OX_i4FBqBI2BZCfaoQ)

# 部署

```shell
kubectl apply -f https://raw.githubusercontent.com/togettoyou/mini-k8s-proxy/master/mini-k8s-proxy.yaml
```

若无法连接 GitHub ，请自行下载 [yaml 文件](https://github.com/togettoyou/mini-k8s-proxy/blob/master/mini-k8s-proxy.yaml)

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
FROM golang:1.16

ARG KUBE_VERSION

ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct

RUN go get k8s.io/code-generator@$KUBE_VERSION; exit 0
RUN go get k8s.io/apimachinery@$KUBE_VERSION; exit 0

RUN mkdir -p $GOPATH/src/k8s.io/{code-generator,apimachinery}
RUN cp -R $GOPATH/pkg/mod/k8s.io/code-generator@$KUBE_VERSION $GOPATH/src/k8s.io/code-generator
RUN cp -R $GOPATH/pkg/mod/k8s.io/apimachinery@$KUBE_VERSION $GOPATH/src/k8s.io/apimachinery
RUN chmod +x $GOPATH/src/k8s.io/code-generator/generate-groups.sh

WORKDIR $GOPATH/src/k8s.io/code-generator

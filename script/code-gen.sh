#!/bin/bash -e

set -e -o pipefail

PROJECT_MODULE="mini-k8s-proxy"
IMAGE_NAME="kubernetes-codegen:latest"

echo "Building codegen Docker image..."
docker build --build-arg KUBE_VERSION=v0.20.2 -f "./script/codegen.Dockerfile" \
            -t "${IMAGE_NAME}" \
            "."

cmd="/go/src/k8s.io/code-generator/generate-groups.sh all \
    ${PROJECT_MODULE}/pkg/generated \
    ${PROJECT_MODULE}/pkg/apis \
    miniproxy:v1alpha1 \
    --go-header-file=/go/src/${PROJECT_MODULE}/script/boilerplate.go.txt"

echo "Generating clientSet code ..."
docker run --rm \
           -v "$(pwd):/go/src/${PROJECT_MODULE}" \
           -w "/go/src/${PROJECT_MODULE}" \
           "${IMAGE_NAME}" $cmd
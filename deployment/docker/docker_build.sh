#!/bin/bash

# go mod 改变更新基础镜像
function buildBase() {
    tmpF="./base_changed"
    localBaseHash=$(md5sum go.mod | awk '{print $1}')_$(md5sum go.sum | awk '{print $1}')
    if [ -f "$tmpF" ]; then
        oldHash=$(cat $tmpF)
        if [[ "${oldHash}" == "${localBaseHash}" ]]; then
            echo "The two strings are the same $oldHash"
            return 0
        fi
    fi

    if docker build -f ./deployment/docker/base.Dockerfile -t "$TAG_NAME_SHA" -t "$TAG_NAME_LATEST" .; then
        echo "$localBaseHash" >$tmpF
        printf "base image build success: %s" "$TAG_NAME_SHA"
    else
       echo "base image build fail!!!"
       exit 1
    fi
}

export DOCKER_BUILDKIT=0
export COMPOSE_DOCKER_CLI_BUILD=0

SHA="$1"
if [ -z "$SHA" ]; then
  SHA="$(git rev-parse --short HEAD)"
fi
DOCKERHUB_USERNAME="$2"
DOCKERHUB_PASSWORD="$3"
DOCKERHUB_NAMESPACE="$4"

IMAGE_NAME="blog_base"

TAG_NAME_SHA="$DOCKERHUB_NAMESPACE/$IMAGE_NAME:$SHA"
TAG_NAME_LATEST="$DOCKERHUB_NAMESPACE/$IMAGE_NAME:latest"

docker login -u "$DOCKERHUB_USERNAME" -p "$DOCKERHUB_PASSWORD"
if [ $? -ne 0 ]; then
    echo "Login to Docker Hub file: $DOCKERHUB_USERNAME"
    exit 1
fi

buildBase "$SHA"

docker push "$TAG_NAME_SHA"

echo "$SHA"


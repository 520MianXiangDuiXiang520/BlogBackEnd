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
    tagName="blog_base:$1"
    tagNameLaster="blog_base:latest"
    if docker build -f ./deployment/docker/base.Dockerfile -t "$tagNameLaster" -t "$tagName" .; then
        echo "$localBaseHash" >$tmpF
        printf "base image build success: %s" "$tagName"
    else
       echo "base image build fail!!!"
       exit 1
    fi
}

function build() {
    docker build -f ./deployment/docker/dockerfile -t "blogapi:$1" -t "blogapi:latest" .
}

SHA="$1"
if [ -z "$SHA" ]; then
  SHA="$(git rev-parse --short HEAD)"
fi
DOCKERHUB_USERNAME="$2"
DOCKERHUB_PASSWORD="$3"

docker login -u "$DOCKERHUB_USERNAME" -p "$DOCKERHUB_PASSWORD"
if [ $? -ne 0 ]; then
    echo "Login to Docker Hub file: $DOCKERHUB_USERNAME"
    exit 1
fi


echo "$SHA"

buildBase "$SHA"
build "$SHA"

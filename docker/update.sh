#!/bin/bash
set -e

for dockerfile in $(ls Dockerfile.*) ; do
    sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur # `date`|g" $dockerfile
    docker build -f $dockerfile \
    --rm -t pacur/pacur-${dockerfile#Dockerfile.} .
    sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur|g" $dockerfile
done

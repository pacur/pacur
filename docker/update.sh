#!/bin/bash
set -e

for dockerfile in $(ls Dockerfile.*) ; do
    sed -i -e "s|go get github.com/m0rf30/pacur.*|go get github.com/m0rf30/pacur # `date`|g" $dockerfile
    docker build -f $dockerfile \
    --rm -t m0rf30/pacur-${dockerfile#Dockerfile.} .
    sed -i -e "s|go get github.com/m0rf30/pacur.*|go get github.com/m0rf30/pacur|g" $dockerfile
done

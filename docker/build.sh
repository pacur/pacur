#!/bin/bash
set -e

for dockerfile in $(ls Dockerfile.*) ; do
    docker build -f $dockerfile \
    --rm -t m0rf30/pacur-${dockerfile#Dockerfile.} .
done

#!/bin/bash
set -e

for dockerfile in $(ls Dockerfile.*) ; do
    docker build -f $dockerfile \
    --rm -t pacur/pacur-${dockerfile#Dockerfile.} .
done

#!/bin/bash
set -e
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

for dockerfile in $(ls Dockerfile.*) ; do
    docker push m0rf30/pacur-${dockerfile#Dockerfile.}
done

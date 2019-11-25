#!/bin/bash
set -e
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

for dockerfile in $(ls Dockerfile.*) ; do
    docker push pacur/pacur-${dockerfile#Dockerfile.}
done

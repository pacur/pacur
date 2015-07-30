#!/bin/bash
cd "$( dirname "${BASH_SOURCE[0]}" )"

cd centos-6
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur #`date`|g" Dockerfile
docker build --rm -t centos-6 .
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur|g" Dockerfile
cd ..

cd centos-7
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur #`date`|g" Dockerfile
docker build --rm -t centos-7 .
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur|g" Dockerfile
cd ..

cd debian-jessie
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur #`date`|g" Dockerfile
docker build --rm -t debian-jessie .
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur|g" Dockerfile
cd ..

cd debian-wheezy
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur #`date`|g" Dockerfile
docker build --rm -t debian-wheezy .
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur|g" Dockerfile
cd ..

cd ubuntu-precise
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur #`date`|g" Dockerfile
docker build --rm -t ubuntu-precise .
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur|g" Dockerfile
cd ..

cd ubuntu-trusty
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur #`date`|g" Dockerfile
docker build --rm -t ubuntu-trusty .
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur|g" Dockerfile
cd ..

cd ubuntu-vivid
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur #`date`|g" Dockerfile
docker build --rm -t ubuntu-vivid .
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur|g" Dockerfile
cd ..

cd ubuntu-wily
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur #`date`|g" Dockerfile
docker build --rm -t ubuntu-wily .
sed -i -e "s|go get github.com/pacur/pacur.*|go get github.com/pacur/pacur|g" Dockerfile
cd ..

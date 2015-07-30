#!/bin/bash
cd "$( dirname "${BASH_SOURCE[0]}" )"

cd centos-6
docker build --rm -t centos-6 .
cd ..

cd centos-7
docker build --rm -t centos-7 .
cd ..

cd debian-jessie
docker build --rm -t debian-jessie .
cd ..

cd debian-wheezy
docker build --rm -t debian-wheezy .
cd ..

cd ubuntu-precise
docker build --rm -t ubuntu-precise .
cd ..

cd ubuntu-trusty
docker build --rm -t ubuntu-trusty .
cd ..

cd ubuntu-vivid
docker build --rm -t ubuntu-vivid .
cd ..

cd ubuntu-wily
docker build --rm -t ubuntu-wily .
cd ..

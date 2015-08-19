#!/bin/bash
cd "$( dirname "${BASH_SOURCE[0]}" )"

cd archlinux
docker build --rm --no-cache -t archlinux .
cd ..

cd centos-6
docker build --rm --no-cache -t centos-6 .
cd ..

cd centos-7
docker build --rm --no-cache -t centos-7 .
cd ..

cd debian-jessie
docker build --rm --no-cache -t debian-jessie .
cd ..

cd debian-wheezy
docker build --rm --no-cache -t debian-wheezy .
cd ..

cd ubuntu-precise
docker build --rm --no-cache -t ubuntu-precise .
cd ..

cd ubuntu-trusty
docker build --rm --no-cache -t ubuntu-trusty .
cd ..

cd ubuntu-vivid
docker build --rm --no-cache -t ubuntu-vivid .
cd ..

cd ubuntu-wily
docker build --rm --no-cache -t ubuntu-wily .
cd ..

cd genkey
docker build --rm --no-cache -t genkey .
cd ..

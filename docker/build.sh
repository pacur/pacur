#!/bin/bash
set -e
cd "$( dirname "${BASH_SOURCE[0]}" )"

for dir in */ ; do
    cd $dir
    docker build --rm -t pacur/${dir::-1} .
    cd ..
done

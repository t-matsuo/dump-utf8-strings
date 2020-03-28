#!/bin/bash

IMAGE="localhost/dump-utf8-strings-build:latest"
docker build -t $IMAGE . && docker run -d --rm -v `pwd`:/build --name dump-utf8-strings $IMAGE

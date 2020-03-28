#!/bin/bash

if [ "${1:0:1}" != '-' ]; then
    exec $1
fi

go build -a -ldflags="-w -s -extldflags \"-static\"" -gcflags="-trimpath="`pwd` -asmflags=-trimpath=`pwd` -o dump-utf8-strings main.go

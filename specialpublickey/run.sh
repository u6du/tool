#!/usr/bin/env bash

_dirname=$(cd "$(dirname "$0")"; pwd)

cd $_dirname
 
nohup docker run --rm \
-v /root/go:/go \
-v $_dirname:$_dirname -w $_dirname u6du/cloudflare-dns go run main.go &

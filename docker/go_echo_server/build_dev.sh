#!/usr/bin/env bash

CODE_ROOT=$(dirname $(realpath $0))/../..

docker build \
    --rm \
    -t go-echo-server-dev:latest \
    -f $CODE_ROOT/docker/go_echo_server/Dockerfile \
    --build-arg GO_PKG=$(cat $CODE_ROOT/go/go.mod | head -n 1 | awk '{ print $2 }') \
    --build-arg GIT_DESCRIBE=$(git describe --all --always --dirty --long) \
    $CODE_ROOT

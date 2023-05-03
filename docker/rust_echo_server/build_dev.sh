#!/usr/bin/env bash

CODE_ROOT=$(dirname $(realpath $0))/../..

docker build \
    --rm \
    -t rust-echo-server-dev:latest \
    -f $CODE_ROOT/docker/rust_echo_server/Dockerfile \
    --build-arg GIT_DESCRIBE=$(git describe --all --always --dirty --long) \
    $CODE_ROOT

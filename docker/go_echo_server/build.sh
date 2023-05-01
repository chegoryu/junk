#!/usr/bin/env bash

CODE_ROOT=$(dirname $(realpath $0))/../..

docker build --rm -t go-echo-server-dev:latest -f $CODE_ROOT/docker/go_echo_server/Dockerfile $CODE_ROOT

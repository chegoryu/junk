#!/usr/bin/env bash

CODE_ROOT=$(dirname $(realpath $0))/../..

docker-compose -f  $CODE_ROOT/docker/rust_echo_server/docker-compose-dev.yaml up

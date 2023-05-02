#!/usr/bin/env bash

CODE_ROOT=$(dirname $(realpath $0))/../..

docker-compose -f  $CODE_ROOT/docker/go_echo_server/docker-compose-dev.yaml up

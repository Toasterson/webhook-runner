#!/usr/bin/bash

set -ex

## Variables
GOPATH=${GOPATH:-"$(pwd)/usr"}

go get $1
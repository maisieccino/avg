#!/usr/bin/env bash

set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
PATH=$GOPATH/bin:$PATH

which protoc
if [ "$?" -ne "0" ]; then
    echo "Error: you must install protobuf" >&2
    echo "e.g. apk install protobuf" >&2
    exit 1
fi

which protoc-gen-go
if [ "$?" -ne "0" ]; then
    echo "Error: you must install protoc Go bindings." >&2
    echo "Run go get -u github.com/golang/protobuf/protoc-gen-go" >&2
    exit 1
fi

protoc -I pkg/pb pkg/pb/*.proto --go_out=plugins=grpc:pkg/pb
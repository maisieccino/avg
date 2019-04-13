#!/bin/sh -xe
export GO111MODULE=on
export GOOS="$os"
export GOARCH="$arch"
cd src
apk add --no-cache ca-certificates git protobuf bash
GO111MODULE=off CGO_ENABLED=0 go get -u github.com/golang/protobuf/protoc-gen-go           
go mod download
./hack/verify_codegen.sh
version=$(cat ../$version_file)
CGO_ENABLED=0 go build -o "output/avg-${version}-${os}" cmd/client/main.go


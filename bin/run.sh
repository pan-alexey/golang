#!/bin/bash

export GO111MODULE=on
export PATH="$PATH:$(go env GOPATH)/bin"

mkdir -p contract

protoc --proto_path=./proto \
  --go_out=. \
  --go_opt=Mserver.proto=./contract/server \
  ./proto/server.proto

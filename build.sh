#!/usr/bin/env bash

go fmt ./cmd/...
go vet ./cmd/...
if [ $? -ne 0 ] ; then exit; fi

VERSION=$(cat VERSION)
COMMIT=$(git rev-parse --short HEAD)
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
COMPILER="$(go version)"

set -e
export GOFLAGS="-mod=vendor"
go build \
  -ldflags="-s -w -X 'main.version=${VERSION}' -X 'main.commit=${COMMIT}' -X 'main.buildTime=${BUILD_TIME}' -X 'main.compiler=${COMPILER}'" \
  -o bin/financelime-functional-tests cmd/main.go
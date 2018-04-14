#!/usr/bin/env bash

set -e
echo "" > coverage/coverage.out

for d in $(go list ./... | grep -v vendor); do
    go test -race -v -coverprofile=coverage.chunk.out -covermode=atomic $d
    if [ -f coverage.chunk.out ]; then
        cat coverage.chunk.out >> coverage/coverage.out
        rm coverage.chunk.out
    fi
done

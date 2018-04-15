#!/usr/bin/env bash

set -e

DIRECTORY="coverage"

if [ ! -d $DIRECTORY ]; then
    mkdir $DIRECTORY
fi

echo "mode: set" > coverage/coverage.out

for d in $(go list ./... | grep -v vendor); do
    go test -race -v -coverprofile=coverage.chunk.out -covermode=atomic $d
    if [ -f coverage.chunk.out ]; then
        grep -h -v "^mode:" coverage.chunk.out >> coverage/coverage.out
        rm coverage.chunk.out
    fi
done

cat coverage/coverage.out > c.out
rm -rf $DIRECTORY

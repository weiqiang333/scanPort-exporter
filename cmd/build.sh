#!/usr/bin/env bash

export GOARCH=amd64
export GOOS=linux
export GCCGO=gc

version=$1

go build -o scanPort-exporter main.go
chmod +x scanPort-exporter

if [ -z $version ]; then
    version=v0.1
fi

tar -zcvf scanPort-exporter-linux-amd64-${version}.tar.gz \
  scanPort-exporter config/ README.md

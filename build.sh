#!/bin/sh

apk add --update upx git

VERSION=`git fetch --tags && git describe --tags --candidates=1 --dirty --always`
FLAGS=-"s -w -X main.Version=$VERSION"
BIN=build/s3grabber-`uname -s`-`uname -m`-$VERSION

go get ./...
go build -o $BIN -ldflags="$FLAGS"
upx -9 $BIN

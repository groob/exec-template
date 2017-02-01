#!/bin/bash

VERSION="$(git describe --tags --always --dirty)"
NAME="exec-template"

echo "Building $NAME version $VERSION"

mkdir -p build

build() {
  echo -n "=> $1-$2: "
  GOOS=$1 GOARCH=$2 CGO_ENABLED=0 go build -o build/$NAME-$1-$2
  du -h build/$NAME-$1-$2
}

build "darwin" "amd64"
build "linux" "amd64"

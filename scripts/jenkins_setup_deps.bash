#!/bin/bash
set +x
set -e

source ./scripts/jenkins_env.bash

mkdir -p ./goworkspace/bin
mkdir -p ./goworkspace/src/github.com/fuserobotics
ln -fs $(pwd) ./goworkspace/src/github.com/fuserobotics/goterram

pushd ${GOPATH}/src/github.com/fuserobotics/goterram
glide install
go get -u -v github.com/gopherjs/gopherjs
popd

# NPM deps
enable-npm-proxy || true
NODE_ENV='dev' npm install
npm prune
set -x

#!/bin/bash
set +x
set -e

source ./scripts/jenkins_env.bash

mkdir -p ./goworkspace/bin
mkdir -p ./goworkspace/src/github.com/paralin
ln -fs $(pwd) ./goworkspace/src/github.com/paralin/goterram

go get -u -v github.com/gopherjs/gopherjs
pushd ./goworkspace/src/github.com/paralin/goterram
go get -v ./js/
popd

# NPM deps
enable-npm-proxy || true
NODE_ENV='dev' npm install
npm prune
set -x

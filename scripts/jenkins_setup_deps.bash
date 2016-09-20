#!/bin/bash

. ./scripts/jenkins_env.bash

git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

mkdir -p ./goworkspace/bin
mkdir -p ./goworkspace/src/github.com/fuserobotics
ln -fs $(pwd) ./goworkspace/src/github.com/fuserobotics/goterram
# ln -fs $(pwd)/vendor ./goworkspace/src

pushd ${GOPATH}/src/github.com/fuserobotics/goterram
glide install
go get -u -v github.com/gopherjs/gopherjs
popd

# NPM deps
# npm install --unsafe-perm -g semantic-release-cli
pushd js
NODE_ENV='dev' npm install
popd

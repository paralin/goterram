#!/bin/bash

set +x
source ~/.nvm/nvm.sh
source ~/.gvm/scripts/gvm

export PATH="$PATH:${GOPATH}/bin"

export GOPATH=$(pwd)/goworkspace
export PATH="$PATH:${GOPATH}/bin"

set -x

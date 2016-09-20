#!/bin/bash

set +x
source ~/.nvm/nvm.sh
source ~/.gvm/scripts/gvm

export PATH="$PATH:${GOPATH}/bin"

export GOPATH=$(pwd)/goworkspace
export PATH="$PATH:${GOPATH}/bin"

if ./scripts/is_branch.bash release ; then
  export JENKINS_SHOULD_RELEASE=ismaster
else
  unset JENKINS_SHOULD_RELEASE
fi
set -x

#!/bin/bash
set -e
. ~/.bashrc
export GOPATH=$(pwd)/goworkspace
export PATH=$PATH:${GOPATH}/bin
if ./scripts/is_master.bash; then
  export JENKINS_SHOULD_RELEASE=ismaster
else
  unset JENKINS_SHOULD_RELEASE
fi

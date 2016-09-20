#!/bin/bash
set -e
. ~/.bashrc
export GOPATH=$(pwd)/goworkspace
export PATH=$PATH:${GOPATH}/bin

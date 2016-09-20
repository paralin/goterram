#!/bin/bash
set -e

pushd ./js
gopherjs build --output=goterram.js
popd

#!/bin/bash
set -e

pushd ./js
gopherjs build --output=goterram.js
sed -i -e "/global.require = require/d" ./goterram.js
popd

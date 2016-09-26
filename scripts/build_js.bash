#!/bin/bash
set -e

go build -v -o ./.tmp.bin github.com/fuserobotics/goterram/js || true
rm ./.tmp.bin || true
pushd ./js
gopherjs build --output=goterram.js
sed -i -e "/global.require = require/d" ./goterram.js
popd
node ./scripts/build_source_tree.js

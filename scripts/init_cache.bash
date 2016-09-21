#!/bin/bash
set -e

if [ -z "$CACHE_CONTEXT" ]; then
  echo "Set CACHE_CONTEXT and try again."
  exit 1
fi

init-gce-creds
if ! download-cache $CACHE_CONTEXT; then
  echo "Downloading cache failed, starting with empty node_modules."
  mkdir -p node_modules/
else
  echo "Extracting cache..."
  tar -zxf build-cache.tar.gz
  rm build-cache.tar.gz
fi

echo "Recording hashes..."
record-cache ./node_modules/

#!/bin/bash
set -e -o pipefail
./scripts/commit_branches.bash | grep '^refs/heads/release$'


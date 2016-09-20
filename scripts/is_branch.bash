#!/bin/bash
set -e -o pipefail

BRANCHES=$(./scripts/commit_branches.bash)

echo "Branches this commit is in:"
echo $BRANCHES

./scripts/commit_branches.bash | grep "^refs/remotes/origin/$1\$"

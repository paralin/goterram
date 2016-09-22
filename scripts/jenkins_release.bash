#!/bin/bash
set -e

if ./scripts/is_branch_head.bash release ; then
  export JENKINS_SHOULD_RELEASE=ismaster
else
  echo "Not on release branch head, not releasing."
  exit 0
fi

set -x
git remote -v
git fetch origin --tags
git fetch origin release --tags
git branch -D release || true
git checkout -b release origin/release

npm run semantic-release || true

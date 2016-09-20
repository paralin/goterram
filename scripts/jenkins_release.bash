#!/bin/bash
set -e

if ./scripts/is_branch_head.bash release ; then
  export JENKINS_SHOULD_RELEASE=ismaster
else
  echo "Not on release branch head, not releasing."
  exit 0
fi

git checkout -f refs/remotes/origin/release

pushd js
npm run semantic-release || true
popd

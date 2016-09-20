COMMIT=$(git rev-parse HEAD) # expands hash if needed
for BRANCH in $(git for-each-ref --format "%(refname)" refs/heads); do
  if $(git rev-list $BRANCH | fgrep -q $COMMIT); then
    echo $BRANCH
  fi
done

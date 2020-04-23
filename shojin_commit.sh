#!/bin/sh

# commit comment
today=$(date "+%Y%m%d")

# commit and push
git add PastProblems/
git commit -m "精進${today}"
git push

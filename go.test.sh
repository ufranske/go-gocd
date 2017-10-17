#!/usr/bin/env bash

set -e
echo "" > coverage.out

for d in $(glide novendor); do
    go test -v -coverprofile=profile.txt -covermode=atomic $d
    if [ -f profile.txt ]; then
        cat profile.txt >> coverage.out
        rm profile.txt
    fi
done
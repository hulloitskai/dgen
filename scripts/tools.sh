#!/usr/bin/env bash

set -e  # exit on failure

## Install golint.
mkdir -p "$GOBIN"
echo "Contents of $GOBIN:"
ls -l "$GOBIN" || true

if ! command -v golint > /dev/null; then
  rm -rf "${GOBIN}/golint"
  echo "Installing 'golint'..."
  GO111MODULE=off go get -u golang.org/x/lint/golint
fi
echo "golint: $(command -v golint)"


## Install clipboard program, if on Linux.
if [ "$TRAVIS_OS_NAME" == "linux" ]; then
  sudo apt-get update;
  sudo apt-get install xclip;
fi
echo "xclip: $(command -v xclip)"

set +e

#!/usr/bin/env bash

# because packagr/bumpr depends on compiled binaries (even for testing) we'll be building the test binaries first in the "build container" and then
# executing them in a "runtime container" to get coverage/profiling data.
#
# this script generates the test binaries in the "build container"

set -e

for d in $(go list ./...); do
    # determine the output path
    OUTPUT_PATH=$(echo "$d" | sed -e "s/^github.com\/packagrio\/bumpr\///")
    echo "Generating TEST BINARY: ${OUTPUT_PATH}/test_binary_${1}"
    mkdir -p /caches/test-binaries/${OUTPUT_PATH}
    go test -mod vendor -race -covermode=atomic -tags="static $1" -c -o=/caches/test-binaries/${OUTPUT_PATH}/test_binary_${1} $d

    if [ -d "${d}/testdata" ];
    then
      # copy the testdata directory for this binary if present.
      mkdir -p /caches/test-binaries/${OUTPUT_PATH}/testdata
      cp -r "${d}/testdata/." /caches/test-binaries/${OUTPUT_PATH}/testdata
    fi

done
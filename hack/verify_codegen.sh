#!/usr/bin/env bash

set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
PATH=$GOPATH/bin:$PATH

DIFFROOT="${SCRIPT_ROOT}/pkg/pb"
TMP_DIFFROOT="${SCRIPT_ROOT}/_tmp/pkg/pb"
_tmp="${SCRIPT_ROOT}/_tmp"

cleanup() {
    rm -rf "${_tmp}"
}
trap "cleanup" EXIT SIGINT

cleanup

which protoc
if [ "$?" -ne "0" ]; then
    echo "Error: you must install protobuf" >&2
    echo "e.g. apk install protobuf" >&2
    exit 1
fi

which protoc-gen-go
if [ "$?" -ne "0" ]; then
    echo "Error: you must install protoc Go bindings." >&2
    echo "Run go get -u github.com/golang/protobuf/protoc-gen-go" >&2
    exit 1
fi

mkdir -p "${TMP_DIFFROOT}"
cp -a "${DIFFROOT}"/* "${TMP_DIFFROOT}"

"${SCRIPT_ROOT}/hack/update_codegen.sh"
echo "diffing ${DIFFROOT} against freshly generated codegen"
ret=0
diff -Naupr "${DIFFROOT}" "${TMP_DIFFROOT}" || ret=$?
cp -a "${TMP_DIFFROOT}"/* "${DIFFROOT}"
if [[ $ret -eq 0 ]]; then
    echo "${DIFFROOT} up to date"
else
    echo "${DIFFROOT} is out of date. Please run hack/update_codegen.sh"
    exit 1
fi
#!/usr/bin/env bash

set -e

APP="redisbench"
PLATFORMS="darwin freebsd linux windows"
ARCHS="amd64"

for OS in ${PLATFORMS[@]}; do 
    for ARCH in ${ARCHS[@]}; do 
        NAME="${APP}-${OS}-${ARCH}"
        if [[ "${OS}" == "windows" ]]; then
            NAME="${NAME}.exe"
        fi

        env GOOS=${OS} GOARCH=${ARCH} go build -o release/${NAME} .
    done
done
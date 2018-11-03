#!/usr/bin/env bash

set -e

cd "${0%/*}"

source .version
NEW_VERSION="$(./semver.sh bump minor ${VERSION})"

echo "bumping version: ${VERSION} => ${NEW_VERSION}"

GOOS=linux GOARCH=arm vgo build -a -installsuffix cgo -o ../.bin/bin ../../../cmd/server/main.go
docker build -t bobrnor/hl-course:${NEW_VERSION} ../.

source docker.io.auth

docker login -u ${USERNAME} -p ${PASSWORD}
docker push bobrnor/hl-course:${NEW_VERSION}
docker logout

echo "VERSION=${NEW_VERSION}" > .version
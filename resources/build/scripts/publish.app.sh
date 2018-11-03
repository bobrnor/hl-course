#!/usr/bin/env bash

set -e

cd "${0%/*}"

GOOS=linux GOARCH=arm vgo build -a -installsuffix cgo -o ../.bin/bin ../../../cmd/server/main.go
docker build -t bobrnor/hl-course ../.

source docker.io.auth

docker login -u $USERNAME -p $PASSWORD
docker push bobrnor/hl-course
docker logout
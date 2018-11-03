#!/usr/bin/env bash

cd "${0%/*}"

source .version
helm upgrade --install hl-course-server-release --namespace hl-course-ns --set "app.version=${VERSION}" ../helm/
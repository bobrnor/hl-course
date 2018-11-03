#!/usr/bin/env bash

cd "${0%/*}"

helm upgrade --install hl-course-server-release --namespace hl-course-ns ../helm/
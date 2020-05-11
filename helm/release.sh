#!/usr/bin/env bash

set -x
set -e
set -o nounset
set -o pipefail

helm package helm/
mv *.tgz docs/
helm repo index docs/ --url https://github.io/fielmann-ag/version-monitor/docs

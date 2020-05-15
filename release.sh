#!/usr/bin/env bash

set -e
set -o nounset
set -o pipefail

URL=${URL:-'https://raw.githubusercontent.com/fielmann-ag/version-monitor/master/docs'}

pushd docs >/dev/null
helm package ../helm/
helm repo index . --url "$URL"
popd >/dev/null

#!/usr/bin/env sh

dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd -P)
target=ae-ops-infrastructure

if ! fly status -t "$target"; then
  fly login \
    --target "$target" \
    --team-name "$target" \
    --concourse-url=https://ci.mgmt.ae.cloudhh.de/
fi

fly -t "$target" set-pipeline \
  --config="$dir/pipeline.yaml" --pipeline=version-monitor \
  --check-creds

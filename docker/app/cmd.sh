#!/bin/bash

readonly APP_ENTRY_POINT="/workspace/src/main.go"
readonly DLV_API_VERSION=2
readonly DLV_OUTPUT="/root/.cache/go-build/dlv"
readonly DLV_PORT=":40000"

if [[ -n "${DEBUG}" ]] && [[ "${DEBUG}" = "true" ]]; then
  dlv debug \
    --headless \
    --listen="${DLV_PORT}" \
    --api-version="${DLV_API_VERSION}" \
    --accept-multiclient \
    --output "${DLV_OUTPUT}" \
    "${APP_ENTRY_POINT}"
else
  go run "${APP_ENTRY_POINT}"
fi

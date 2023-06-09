#!/usr/bin/env bash
set -eu

readonly APP_NAME="${1}"
readonly SCRIPT_DIR="$(dirname "$0")"

. "${SCRIPT_DIR}/config"

atlas migrate lint \
  --dir="file://app/${APP_NAME}/ent/migrations" \
  --dev-url="mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}" \
  --latest=1

#!/usr/bin/env bash
set -eu

readonly APP_NAME="${1}"

readonly SCRIPT_DIR="$(dirname "$0")"
. "${SCRIPT_DIR}/config.sh"

atlas migrate apply \
  --dir "file://app/${APP_NAME}/ent/migrations" \
  --url "mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${APP_NAME}"

atlas migrate apply \
  --dir "file://app/${APP_NAME}/ent/migrations" \
  --url "mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${APP_NAME}_test"

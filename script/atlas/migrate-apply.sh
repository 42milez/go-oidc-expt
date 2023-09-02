#!/usr/bin/env bash
set -eu

readonly DBNAME="idp"

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
. "${SCRIPT_DIR}/config.sh"

atlas migrate apply \
  --dir "file://app/ent/migrations" \
  --url "mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DBNAME}"

atlas migrate apply \
  --dir "file://app/${APP_NAME}/ent/migrations" \
  --url "mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DBNAME}_test"

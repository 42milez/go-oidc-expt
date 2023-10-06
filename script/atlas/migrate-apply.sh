#!/usr/bin/env bash
set -eu

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
. "${SCRIPT_DIR}/config.sh"

readonly DB_NAMES=(${1//,/ })
readonly MIGRATION_DIR="file://app/ent/migrations"

function Migrate() {
  local -r dbname="${1}"
  local -r dir="${MIGRATION_DIR}"
  local -r url="mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${dbname}"
  echo "apply '${dir}' to '${url}'"
  atlas migrate apply --dir "${dir}" --url "${url}"
}

for dbname in "${DB_NAMES[@]}"; do
  Migrate "${dbname}"
done

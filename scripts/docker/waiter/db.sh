#!/usr/bin/env bash
set -e

function wait_db() {
  local -r SERVICE="${1}"
  local -r DB1_HOST="${2}"
  local -r DB1_PORT="${3}"
  local -r DB_USER="${4}"
  local -r HEALTHCHECK_COMMAND="mysql -h '${DB1_HOST}' -P '${DB1_PORT}' -u '${DB_USER}' -e 'SELECT 1;'"

  wait_service "${SERVICE}" "${HEALTHCHECK_COMMAND}"
}

wait_db "db1" "127.0.0.1" 13306 "root"
wait_db "db2" "127.0.0.1" 23306 "root"

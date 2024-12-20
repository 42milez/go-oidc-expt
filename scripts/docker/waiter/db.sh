#!/usr/bin/env bash
set -e

function wait_db() {
  local -r SERVICE="${1}"
  local -r DB_HOST="${2}"
  local -r DB_PORT="${3}"
  local -r DB_USER="${4}"
  local -r HEALTHCHECK_COMMAND="mysql -h '${DB_HOST}' -P '${DB_PORT}' -u '${DB_USER}' -e 'SELECT 1;'"
  local -r SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

  . "${SCRIPT_DIR}/waiter.sh"

  wait_service "${SERVICE}" "${HEALTHCHECK_COMMAND}"
}

wait_db "$@"

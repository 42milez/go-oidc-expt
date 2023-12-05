#!/usr/bin/env bash
set -e

function wait_app() {
  local -r SERVICE="${1}"
  local -r HOST="${2}"
  local -r PORT="${3}"
  local -r HEALTHCHECK_COMMAND="curl -f http://${HOST}:${PORT}/health"
  local -r SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

  . "${SCRIPT_DIR}/waiter.sh"

  wait_service "${SERVICE}" "${HEALTHCHECK_COMMAND}"
}

wait_app "$@"

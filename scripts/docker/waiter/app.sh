#!/usr/bin/env bash
set -e

function wait_app() {
  local -r SERVICE="${1}"
  local -r HOST="${2}"
  local -r PORT="${3}"
  local -r HEALTHCHECK_COMMAND="curl -f http://${HOST}:${PORT}/health"

  wait_service "${SERVICE}" "${HEALTHCHECK_COMMAND}"
}

wait_app "app" "localhost" 8080
wait_app "app-ci" "localhost" 8081

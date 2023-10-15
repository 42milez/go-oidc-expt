#!/usr/bin/env bash
set -e

if [[ -z "${SCRIPT_DIR}" ]]; then
  readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
  . "${SCRIPT_DIR}/waiter.sh"
fi

readonly APP_HEALTHCHECK_COMMAND="curl -f http://localhost:8080/health"
readonly APP_CI_HEALTHCHECK_COMMAND="curl -f http://localhost:8081/health"

WaitService "app" "${APP_HEALTHCHECK_COMMAND}"
WaitService "app-ci" "${APP_CI_HEALTHCHECK_COMMAND}"

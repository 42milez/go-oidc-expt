#!/usr/bin/env bash
set -e

readonly APP_SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
. "${APP_SCRIPT_DIR}/waiter.sh"

readonly APP_HEALTHCHECK_COMMAND="curl -f http://localhost:8080/health"
readonly APP_CI_HEALTHCHECK_COMMAND="curl -f http://localhost:8081/health"

WaitService "app" "${APP_HEALTHCHECK_COMMAND}"
WaitService "app-ci" "${APP_CI_HEALTHCHECK_COMMAND}"

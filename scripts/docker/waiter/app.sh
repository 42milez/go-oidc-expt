#!/usr/bin/env bash
set -eu

readonly APP_HEALTHCHECK_COMMAND="curl -f http://localhost:8080/health"
readonly APP_CI_HEALTHCHECK_COMMAND="curl -f http://localhost:8081/health"

WaitService "app" "${APP_HEALTHCHECK_COMMAND}"
WaitService "app-ci" "${APP_CI_HEALTHCHECK_COMMAND}"

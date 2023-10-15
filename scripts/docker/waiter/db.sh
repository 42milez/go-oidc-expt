#!/usr/bin/env bash
set -eu

readonly DB_HOST="127.0.0.1"
readonly DB_PORT=13306
readonly DB_USER="root"
readonly HEALTHCHECK_COMMAND="mysql -h '${DB_HOST}' -P '${DB_PORT}' -u '${DB_USER}' -e 'SELECT 1;'"

WaitService "db" "${HEALTHCHECK_COMMAND}"

#!/usr/bin/env bash
#shellcheck disable=SC2155
set -eu

function up() {
  local -r SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

  docker-compose up -d app app-ci cache db1 db2 log

  . "${SCRIPT_DIR}/waiter/waiter.sh"
  . "${SCRIPT_DIR}/waiter/app.sh"
  . "${SCRIPT_DIR}/waiter/db.sh"
}

up

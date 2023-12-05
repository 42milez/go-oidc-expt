#!/usr/bin/env bash
#shellcheck disable=SC2155
set -eu

function up() {
  local -r SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

  docker-compose up -d app app-ci cache db1 db2 log

  . "${SCRIPT_DIR}/waiter/app.sh" "app" "localhost" 8080
  . "${SCRIPT_DIR}/waiter/app.sh" "app-ci" "localhost" 8081
  . "${SCRIPT_DIR}/waiter/db.sh" "db1" "127.0.0.1" 13306 "root"
  . "${SCRIPT_DIR}/waiter/db.sh" "db2" "127.0.0.1" 23306 "root"
}

up

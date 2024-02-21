#!/usr/bin/env bash
#shellcheck disable=SC2155
set -eu

function up() {
  local -r SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

  docker-compose up -d app app-ci cache db log

  . "${SCRIPT_DIR}/waiter/app.sh" "app" "localhost" 8080
  . "${SCRIPT_DIR}/waiter/app.sh" "app-ci" "localhost" 8081
  . "${SCRIPT_DIR}/waiter/db.sh" "db" "127.0.0.1" 13306 "root"
}

up

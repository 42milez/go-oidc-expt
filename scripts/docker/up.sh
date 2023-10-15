#!/usr/bin/env bash
set -eu

readonly UP_SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
. "${UP_SCRIPT_DIR}/waiter/waiter.sh"

docker-compose up -d app cache db log app-ci

. "${UP_SCRIPT_DIR}/waiter/app.sh"
. "${UP_SCRIPT_DIR}/waiter/db.sh"

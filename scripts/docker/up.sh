#!/usr/bin/env bash
set -eu

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
. "${SCRIPT_DIR}/waiter/waiter.sh"

docker-compose up -d app cache db log app-ci

. "${SCRIPT_DIR}/waiter/app.sh"
. "${SCRIPT_DIR}/waiter/db.sh"

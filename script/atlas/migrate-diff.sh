#!/usr/bin/env bash
set -eu

readonly APP_NAME="${1}"
readonly MIGRATION_NAME="${2}"

readonly SCRIPT_DIR="$(dirname "$0")"
. "${SCRIPT_DIR}/config.sh"

go run -mod=mod "app/${APP_NAME}/ent/cmd/migrate/main.go" "${MIGRATION_NAME}"

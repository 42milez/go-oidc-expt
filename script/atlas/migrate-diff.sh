#!/usr/bin/env bash
set -eu

readonly MIGRATION_NAME="${1}"

readonly SCRIPT_DIR="$(dirname "$0")"
. "${SCRIPT_DIR}/config.sh"

go run -mod=mod "app/ent/cmd/migrate/main.go" "${MIGRATION_NAME}"

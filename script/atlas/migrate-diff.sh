#!/usr/bin/env bash
set -eu

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
. "${SCRIPT_DIR}/config.sh"

readonly MIGRATION_NAME="${1}"

go run -mod=mod "app/ent/cmd/migrate/diff.go" "${MIGRATION_NAME}"

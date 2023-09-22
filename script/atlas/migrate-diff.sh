#!/usr/bin/env bash
set -eu

readonly MIGRATION_NAME="${1}"

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
. "${SCRIPT_DIR}/config.sh"

go run -mod=mod "app/ent/cmd/migrate/diff.go" "${MIGRATION_NAME}"

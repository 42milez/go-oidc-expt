#!/usr/bin/env bash
set -e

readonly N_LATEST="${1}"

readonly SCRIPT_DIR="$(dirname "$0")"
. "${SCRIPT_DIR}/config.sh"

cmd="atlas migrate lint --dir file://app/ent/migrations --dev-url mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}"

if [[ -n "${N_LATEST}" ]]; then
  cmd_exec="${cmd} --latest ${N_LATEST}"
else
  cmd_exec="${cmd} --git-base main"
fi

eval "${cmd_exec}"

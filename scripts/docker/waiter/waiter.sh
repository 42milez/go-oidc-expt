#!/usr/bin/env bash
set -eu

export INTERVAL=5
export STARTED_AT=$(date +%s)
export TIMEOUT=300

function WaitService() {
  local -r name="${1}"
  local -r cmd="${2}"

  echo "[${name}] Wait for service to be available"

  while ! eval "${cmd}" >/dev/null 2>&1; do
    now=$(date +%s)
    d=$((now-STARTED_AT))

    if [[ ${d} -gt ${TIMEOUT} ]]; then
      echo "[${name}] Timeout"
      exit 1
    fi

    echo "[${name}] Waiting... ${d}s"

    sleep "${INTERVAL}"
  done

  echo "[${name}] Service is ready"
}

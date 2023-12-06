#!/usr/bin/env bash
set -eu

function wait_service() {
  local -r SERVICE="${1}"
  local -r CMD="${2}"
  local -r INTERVAL=5
  local -r TIMEOUT=300
  local -r STARTED_AT=$(date +%s)

  echo "[${SERVICE}] Wait for service to be available (Timeout: ${TIMEOUT}s)"

  while ! eval "${CMD}" >/dev/null 2>&1; do
    now=$(date +%s)
    d=$((now-STARTED_AT))

    if [[ ${d} -gt ${TIMEOUT} ]]; then
      echo "[${SERVICE}] Timeout"
      exit 1
    fi

    echo "[${SERVICE}] Waiting... ${d}s"

    sleep "${INTERVAL}"
  done

  echo "[${SERVICE}] Service is ready"
}

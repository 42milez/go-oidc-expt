#!/usr/bin/env bash
set -eu

readonly DB_HOST="127.0.0.1"
readonly DB_PORT=13306
readonly DB_USER="root"
readonly INTERVAL=5
readonly STARTED_AT=$(date +%s)
readonly TIMEOUT=300

while ! mysql -h "${DB_HOST}" -P "${DB_PORT}" -u "${DB_USER}" -e "SELECT 1;" >/dev/null 2>&1; do
  now=$(date +%s)
  d=$((now-STARTED_AT))

  if [[ ${d} -gt ${TIMEOUT} ]]; then
    echo "timeout: failed to establish database connection"
    exit 1
  fi

  echo "waiting... ${d}s"

  sleep "${INTERVAL}"
done

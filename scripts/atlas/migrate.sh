#!/usr/bin/env bash
#shellcheck disable=SC1001,SC2181,SC2206
set -eu

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

# import .env
for v in $(< "${SCRIPT_DIR}/../../.env"); do
  if [[ "${v}" = ^\s*$ ]] || [[ "${v}" =~ ^# ]]; then
    continue
  fi
  export "${v?}"
done

# Subcommands and its options
#   apply:
#     - s,service: Service name defined in docker-compose.yml.
#     - d,database: Database name.
#   diff:
#     - m,migration-name: Migration filename to be created.
#   lint:
#     - l,latest: The number of latest migration files to be analyzed.
VALID_ARGS=$(getopt -o "s:d:m:l:" --long "service:,database:,migration-name:,latest:" -- "$@")
if [[ $? -ne 0 ]]; then
    exit 1;
fi

eval set -- "$VALID_ARGS"

#shellcheck disable=SC2124
readonly CMD="${@:$#:1}"

# parse arguments
case "${CMD}" in
  apply)
    SERVICE=""
    DATABASE=""
    while true; do
      case "$1" in
        -s | --service)
          SERVICE="$2"
          readonly SERVICE
          shift 2
          ;;
        -d | --database)
          DATABASE="$2"
          readonly DATABASE
          shift 2
          ;;
        --) shift;
          break
          ;;
      esac
    done
  ;;
  diff)
    MIGRATION_NAME=""
    while true; do
      case "$1" in
        -m | --migration-name)
          MIGRATION_NAME="$2"
          readonly MIGRATION_NAME
          shift 2
        ;;
        --) shift;
          break
        ;;
      esac
    done
  ;;
  lint)
    LATEST=""
    while true; do
      case "$1" in
        -l | --latest)
          LATEST="$2"
          readonly LATEST
          shift 2
        ;;
        --) shift;
          break
        ;;
      esac
    done
  ;;
  *)
    echo "unsupported command: ${CMD}"
    exit 1
    ;;
esac

function resource_url() {
  local url=""
  case "${SERVICE}:${DATABASE}" in
    db1:idp | db1:idp_test)
      url="mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB1_PORT}/${DATABASE}"
    ;;
    db2:idp)
      #shellcheck disable=SC2153
      url="mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB2_PORT}/${DATABASE}"
    ;;
    *)
      echo "unknown service or database: ${SERVICE}:${DATABASE}"
      exit 1
    ;;
  esac
  echo "${url}"
}

function run() {
  case "${CMD}" in
    apply)
      url=$(resource_url)
      echo "apply '${MIGRATION_DIR}' to '${url}'"
      atlas migrate apply --dir "${MIGRATION_DIR}" --url "${url}"
    ;;
    diff)
      export DB1_PORT
      export DB_NAME="atlas"
      go run -mod=mod "scripts/ent/migrate/diff/main.go" "${MIGRATION_NAME}"
    ;;
    lint)
      cmd="atlas migrate lint --dir file://app/ent/migrations --dev-url mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB1_PORT}/${DB_NAME}"
      if [[ -n "${LATEST}" ]]; then
        cmd_exec="${cmd} --latest ${LATEST}"
      else
        cmd_exec="${cmd} --git-base main"
      fi
      eval "${cmd_exec}"
    ;;
  esac
}

run

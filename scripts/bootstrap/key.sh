#!/usr/bin/env bash
set -eu

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

readonly KEY_DIR="app/idp/security/secret/key"
readonly HASH_KEY_FILE="${KEY_DIR}/hash.key"
readonly BLOCK_KEY_FILE="${KEY_DIR}/block.key"
readonly HASH_KEY_LENGTH=64
readonly BLOCK_KEY_LENGTH=32

echo "create directory: ${KEY_DIR}"
mkdir -p "${KEY_DIR}"

echo "create hash key: ${HASH_KEY_FILE}"
go run "${SCRIPT_DIR}/../random_string.go" --len "${HASH_KEY_LENGTH}" > "${HASH_KEY_FILE}"

echo "create block key: ${BLOCK_KEY_FILE}"
go run "${SCRIPT_DIR}/../random_string.go" --len "${BLOCK_KEY_LENGTH}" > "${BLOCK_KEY_FILE}"

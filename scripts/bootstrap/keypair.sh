#!/usr/bin/env bash
set -eu

readonly KEY_PAIR_DIR="app/idp/security/secret/keypair"
readonly PRIVATE_KEY_FILE="${KEY_PAIR_DIR}/private.pem"
readonly PUBLIC_KEY_FILE="${KEY_PAIR_DIR}/public.pem"
readonly CURVE="prime256v1"

echo "create directory: ${KEY_PAIR_DIR}"
mkdir -p "${KEY_PAIR_DIR}"

echo "create private key: ${PRIVATE_KEY_FILE}"
openssl ecparam -genkey -name "${CURVE}" -noout -out "${PRIVATE_KEY_FILE}" &>/dev/null

echo "create public key: ${PUBLIC_KEY_FILE}"
openssl ec -in "${PRIVATE_KEY_FILE}" -pubout -out "${PUBLIC_KEY_FILE}" &>/dev/null

#!/usr/bin/env bash
set -eu

readonly CERT_DIR="app/pkg/xjwt/cert"
readonly CURVE="prime256v1"

mkdir -p "${CERT_DIR}"
openssl ecparam -genkey -name "${CURVE}" -noout -out "${CERT_DIR}/private.pem"
openssl ec -in "${CERT_DIR}/private.pem" -pubout -out "${CERT_DIR}/public.pem"

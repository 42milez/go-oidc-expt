#!/usr/bin/env bash
set -eu

readonly WORK_DIR="${GITHUB_WORKSPACE}"

echo "create directory: ${WORK_DIR}/app/pkg/xjwt/cert"
mkdir -p "${WORK_DIR}/app/pkg/xjwt/cert"

echo "create private key: ${WORK_DIR}/app/pkg/xjwt/cert/private.pem"
openssl ecparam -genkey -name prime256v1 -noout -out "${WORK_DIR}/app/pkg/xjwt/cert/private.pem"

echo "create public key: ${WORK_DIR}/app/pkg/xjwt/cert/public.pem"
openssl ec -in "${WORK_DIR}/app/pkg/xjwt/cert/private.pem" -pubout -out "${WORK_DIR}/app/pkg/xjwt/cert/public.pem"

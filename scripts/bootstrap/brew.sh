#!/usr/bin/env bash
set -eu

readonly FORMULAS=(
  ariga/tap/atlas
  daveshanley/vacuum/vacuum
  golangci-lint
  gnu-getopt
  mkcert
  nss
  openssl@3
  # docker
  # docker-buildx
  # docker-compose
  #lima
)

brew install "${FORMULAS[@]}"

#!/usr/bin/env bash
set -eu

readonly FORMULAS=(
  ariga/tap/atlas
  daveshanley/vacuum/vacuum
  # docker
  # docker-buildx
  # docker-compose
  golangci-lint
  gnu-getopt
  #lima
  openssl@3
)

brew install "${FORMULAS[@]}"

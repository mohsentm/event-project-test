#!/bin/bash

# shellcheck disable=SC2164
cd "${INSTALL_DIR:=/app}"

# Check .env
if [[ ! -f .env ]]; then
  # create .env
  cp .env.example .env
fi


go build -o ./bin/api ./cmd/api

./bin/api
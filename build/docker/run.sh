#!/bin/bash

# shellcheck disable=SC2164
cd "${INSTALL_DIR:=/app}"

go build -o ./bin/api ./cmd/api

./bin/api
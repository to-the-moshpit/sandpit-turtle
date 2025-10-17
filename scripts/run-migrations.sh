#!/usr/bin/env bash
. ./scripts/set-env.sh

echo "Running migrations ${DATABASE_CONNECTION_STRING}"
migrate -path ./migrations  -database "${DATABASE_CONNECTION_STRING}" -verbose up
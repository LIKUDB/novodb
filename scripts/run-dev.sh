#!/usr/bin/env bash
# Levanta NovoDB en modo desarrollo con datos en ./data-dev
set -euo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")/.."

export NOVODB_DATA="${NOVODB_DATA:-./data-dev}"
export NOVODB_LOG_LEVEL="${NOVODB_LOG_LEVEL:-debug}"
export NOVODB_FAST_STARTUP="${NOVODB_FAST_STARTUP:-true}"

go run ./cmd/novodb

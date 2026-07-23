#!/usr/bin/env bash
# Compila el binario de NovoDB en ./bin/novodb
set -euo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")/.."

mkdir -p bin
go build -o bin/novodb ./cmd/novodb
echo "Binario generado en bin/novodb"

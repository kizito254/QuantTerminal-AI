#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUT_DIR="$ROOT_DIR/dist"
OUT_FILE="$OUT_DIR/QuantTerminalAI.exe"

mkdir -p "$OUT_DIR"
GOOS=windows GOARCH=amd64 go build -o "$OUT_FILE" "$ROOT_DIR/cmd/quantterminalai"

echo "Built Windows executable: $OUT_FILE"

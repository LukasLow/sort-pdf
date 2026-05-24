#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
VERSION=$(grep '^const AppVersion' "$ROOT/src/version.go" | sed 's/.*"\(.*\)".*/\1/')

if [[ "$OSTYPE" == "darwin"* ]]; then
  sed -i "" "s/\"productVersion\": \".*\"/\"productVersion\": \"$VERSION\"/" "$ROOT/wails.json"
else
  sed -i "s/\"productVersion\": \".*\"/\"productVersion\": \"$VERSION\"/" "$ROOT/wails.json"
fi

echo "wails.json productVersion -> $VERSION"

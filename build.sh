#!/bin/bash
set -euo pipefail

echo "Building Tailwind CSS..."
npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --minify

if command -v go >/dev/null 2>&1; then
  echo "Building Go binary..."
  go build -o bin/server ./cmd/server
else
  echo "Go toolchain not available; skipping binary build (handled by Vercel functions)."
fi

echo "Build complete!"

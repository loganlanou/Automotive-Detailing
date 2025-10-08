#!/bin/bash
set -e

echo "Installing templ..."
go install github.com/a-h/templ/cmd/templ@latest

echo "Generating templ templates..."
templ generate

echo "Building Tailwind CSS..."
npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --minify

echo "Building Go binary..."
go build -o bin/server ./cmd/server

echo "Build complete!"

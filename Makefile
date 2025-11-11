.PHONY: help dev build migrate test fmt clean install templ tailwind

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ## Install dependencies
	go mod download
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/air-verse/air@latest
	npm install -D tailwindcss@next @tailwindcss/cli@next postcss autoprefixer

templ: ## Generate templ files
	templ generate

tailwind: ## Build Tailwind CSS
	npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --minify

tailwind-watch: ## Watch and build Tailwind CSS
	npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --watch

migrate: ## Run database migrations
	@mkdir -p data
	sqlite3 data/detailing.db < pkg/db/schema.sql
	@echo "✓ Database migrated"

sqlc: ## Generate SQLC code
	sqlc generate

dev: templ tailwind ## Run development server with hot reload
	air

build: templ tailwind ## Build production binary
	go build -o bin/server cmd/server/main.go

test: ## Run tests
	go test -v ./...

fmt: ## Format code
	go fmt ./...
	templ fmt .

clean: ## Clean build artifacts
	rm -rf bin/
	rm -rf web/static/css/output.css
	rm -f data/detailing.db

run: ## Run the server
	go run cmd/server/main.go

# Build stage
FROM golang:1.23-alpine AS builder

# Install dependencies
RUN apk add --no-cache nodejs npm git

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

COPY package*.json ./
RUN npm install

# Copy source code
COPY . .

# Build CSS
RUN npm run build:css

# Generate templ templates
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

# Build Go binary
RUN go build -o bin/server ./cmd/server

# Runtime stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary and static files
COPY --from=builder /app/bin/server .
COPY --from=builder /app/web ./web
COPY --from=builder /app/public ./public

# Create data directory
RUN mkdir -p ./data

EXPOSE 8080

CMD ["./server"]

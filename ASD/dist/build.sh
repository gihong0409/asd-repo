#!/bin/bash
set -e

echo "🔨 Building Go application..."

# Go 바이너리 빌드
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -o ../app ../

echo "🐳 Building Docker image..."
docker build -t asd:latest ../

echo "✅ Build completed!"

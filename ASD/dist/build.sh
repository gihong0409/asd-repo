#!/bin/bash
set -e

echo "🔨 Building ASD..."

# Go 모듈 정리
echo "📦 Tidying up Go modules..."
go mod tidy

# 빌드
echo "🏗️  Building binary..."
go build -o asd -ldflags "-X main.buildTime=$(date +'%Y/%m/%d_%H:%M:%S')" ../main.go

echo "✅ Build completed!"
echo "Run with: ./script.sh"


export BENTLEY=false

export BENZ=false

export FERRARI=false

export TESLA=false

export MARS=false

export SATURN=true

./asd
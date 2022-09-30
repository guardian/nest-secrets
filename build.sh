#!/usr/bin/env bash

set -e

mkdir -p bin

# Targetting AMD (x86) EC2 
env GOOS=linux GOARCH=amd64 go build -o nest-secrets-linux-amd64 main.go
mv nest-secrets-linux-amd64 bin/

# Targetting ARM64 (Graviton) EC2 
env GOOS=linux GOARCH=arm64 go build -o nest-secrets-linux-arm64 main.go
mv nest-secrets-linux-arm64 bin/

# Targetting AMD (x86) Mac 
env GOOS=darwin GOARCH=amd64 go build -o nest-secrets-darwin-amd64 main.go
mv nest-secrets-darwin-amd64 bin/

# Targetting ARM64 Mac (Apple silicon) 
env GOOS=darwin GOARCH=arm64 go build -o nest-secrets-darwin-arm64 main.go
mv nest-secrets-darwin-arm64 bin/
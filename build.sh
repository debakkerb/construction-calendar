#!/usr/bin/env bash

echo "Building binary for Windows"

GOOS=windows GOARCH=amd64 go build -o bin/cc-amd64.exe main.go

echo "Building binary for OSX (64-bit)"

GOOS=darwin GOARCH=amd64 go build -o bin/cc-amd64-darwin main.go

echo "Building binary for Linux"

GOOS=linux GOARCH=amd64 go build -o bin/cc-amd64-linux main.go
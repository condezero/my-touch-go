#!/bin/bash

echo "Building..."

## Build for windows
GOOS=windows go build -o builds/win/touch.exe main.go
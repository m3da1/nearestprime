#!/bin/bash

export GIN_MODE=release
echo "Building project"
go build -o nearestprime cmd/appserver/main.go
echo "Running project"
./nearestprime
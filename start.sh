#!/bin/bash

export GIN_MODE=release
go build -o nearestprime cmd/appserver/main.go
./nearestprime
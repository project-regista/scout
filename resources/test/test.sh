#!/bin/bash
set -eux

# Go get package dependencies
go get -d ./...

# Run go fmt
go fmt ./...

# Run unit and benchmark tests
go test ./... -v -bench . -benchmem

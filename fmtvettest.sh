#!/usr/bin/env bash

set -xv

go fmt ./cmd/... ./pkg/...
go vet ./cmd/... ./pkg/...
go test ./cmd/... ./pkg/...

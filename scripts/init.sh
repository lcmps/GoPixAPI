#!/usr/bin/env bash
CGO_ENABLED=0 go build -ldflags="-s -w" -o GoPix  -v .
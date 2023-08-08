#!/bin/bash

set -o errexit

export GOOS=linux

go mod tidy

CGO_ENABLED=0 go build -o k8s_customize_controller  main.go

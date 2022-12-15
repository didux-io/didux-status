#!/usr/bin/env bash

.PHONY: didux-status-linux-amd64


COMPANY=Didux
AUTHOR=go-didux

DIR=$(shell pwd)
PACKAGES=$(shell find ./ -type d -not -path '\./')

SRC_DIR="./"


GOBIN=$(shell pwd)/build/bin
GO ?= 1.13

run:
	go run main.go

build:
	go build -o didux-status main.go

linux:
	env GOOS=linux GOARCH=amd64 go build -o didux-status main.go
	@echo "Linux amd64 cross compilation done"

# ********* END BUILD TASKS *********

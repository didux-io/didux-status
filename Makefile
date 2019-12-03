#!/usr/bin/env bash

.PHONY: didux-status-linux-amd64


COMPANY=Didux
AUTHOR=go-didux

DIR=$(shell pwd)
PACKAGES=$(shell find ./src -type d -not -path '\./src')

SRC_DIR="src/"


GOBIN=$(shell pwd)/build/bin
GO ?= 1.11

run:
	go run main.go

build:
	go build -o didux-status main.go

linux:
	./bin/xgo --go=$(GO) --targets=linux/amd64 -v $(shell pwd)
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(shell pwd)/didux-status-linux-amd64 | grep amd64

# ********* END BUILD TASKS *********

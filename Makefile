#!/usr/bin/env bash

.PHONY: smilo-status-linux-amd64


COMPANY=Smilo
AUTHOR=go-smilo

DIR = $(shell pwd)
PACKAGES = $(shell find ./src -type d -not -path '\./src')

SRC_DIR = "src/"


GOBIN = $(shell pwd)/build/bin
GO ?= 1.11

run:
	go run main.go

build: clean
	go build -o smilo-status main.go

linux:
	./bin/xgo --go=$(GO) --targets=linux/amd64 -v $(shell pwd)
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(shell pwd)/smilo-status-linux-amd64 | grep amd64

# ********* END BUILD TASKS *********

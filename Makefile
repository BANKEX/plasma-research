#!/usr/bin/env make

GO ?= go

all: fmt vet test

fmt:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . 2>&1 | grep -v ^vendor/ | tee /dev/stderr)" || \
		(echo >&2 "+ please format Go code with 'gofmt -s'" && false)

vet:
	@echo "+ $@"
	@go tool vet $(shell ls -1 -d */ | grep -v -e vendor)

test:
	@echo "+ $@"
	${GO} test -tags nocgo $(shell go list ./...)

coverage:
	@echo "+ $@"
	${GO} test -tags nocgo $(shell go list ./...) -coverprofile=coverage.out && ${GO} tool cover -html=coverage.out

contracts:
	@echo "+ $@"
	@$(MAKE) -C src/contracts all

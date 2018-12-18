#!/usr/bin/env make

GO ?= go

TARGETDIR := target

HOSTOS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
HOSTARCH := $(shell uname -m)

GOOS ?= ${HOSTOS}
GOARCH ?= ${HOSTARCH}

GOCMD = ./src/node/cmd

TAGS = nocgo

# Set the execution extension for Windows.
ifeq (${GOOS},windows)
    EXE := .exe
endif

OS_ARCH := $(GOOS)_$(GOARCH)$(EXE)

OPERATOR := ${TARGETDIR}/operator_$(OS_ARCH)
VERIFIER := ${TARGETDIR}/verifier_$(OS_ARCH)

all: fmt vet test build

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

build: build/operator build/verifier

build/operator:
	@echo "+ $@"
	${GO} build -tags "$(TAGS)" -ldflags "$(LDFLAGS)" -o ${OPERATOR} ${GOCMD}/operator

build/verifier:
	@echo "+ $@"
	${GO} build -tags "$(TAGS)" -ldflags "$(LDFLAGS)" -o ${VERIFIER} ${GOCMD}/verifier

docker: docker/operator docker/verifier

docker/operator:
	docker build -t plasma-operator:latest -f ./Dockerfile.operator .

docker/verifier:
	docker build -t plasma-verifier:latest -f ./Dockerfile.operator .

export

# Setup Go Variables
GOPATH := $(shell go env GOPATH)
GOBIN := $(PWD)/bin
ifneq ($(GOPRIVATE),)
	GOPRIVATE:= $(GOPRIVATE),
endif
GOPRIVATE := $(GOPRIVATE)github.com/kouzoh

# Invoke shell with new path to enable access to bin
PATH := $(GOBIN):$(PATH)
SHELL := env PATH=$(PATH) bash

# Setup Project Variables
GIT_REF := $(shell git rev-parse --short=7 HEAD)
VERSION ?= commit-$(GIT_REF)
ifneq ($(CI),true)
	VERSION := $(USER)-$(VERSION)
endif
SERVICE_NAME := $(shell grep "^module" go.mod | sed 's:.*/::')

.PHONY: build
build:
	@CGO_ENABLED=0 go build -o bin/server \
      -ldflags "-X main.version=$(VERSION) -X main.serviceName=$(SERVICE_NAME)" \
      ./cmd/http

.PHONY: dependencies
dependencies:
	@./scripts/dependencies.sh

.PHONY: run
run:
	go run ./cmd/http/*.go

.PHONY: test
test:
	go test $(args) -race -cover ./...

.PHONY: test-v
test-v:
	@make test args='-v'

.PHONY: test-long
test-long:
	@go test -v -race -count=5 -coverpkg=./... -coverprofile=coverage.txt ./...

.PHONY: fmt
fmt:
	@find . -iname "*.go" -not -path "./vendor/**" | xargs gofmt -s -w

.PHONY: lint
lint:
	golangci-lint run $(args) ./...
	go-consistent $(cons_args) ./...

.PHONY: lint-fix
lint-fix:
	@make lint args='--fix -v' cons_args='-v'

.PHONY: coverage
coverage:
	@go test -v -coverpkg=./... -coverprofile=cover.out -race ./...

.PHONY: reviewdog
reviewdog:
	reviewdog -conf=.reviewdog.yml -diff="git diff master"

.PHONY: reviewdog-pr
reviewdog-pr:
	golangci-lint run --out-format checkstyle | reviewdog -f=checkstyle -reporter=github-pr-review

.PHONY: changelog
changelog:
	git-chglog --output CHANGELOG.md

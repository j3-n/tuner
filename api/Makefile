SWAG ?= swag
GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
GOMODULES := $(shell go list ./...)
SWAGMODULES := $(shell ls -d internal/*/)

.phony: all
all: swag
    $(GO) run cmd/httpd/main.go

.phony: build
build: swag
    $(GO) build -o build/program/app main.go

.phony: clean
clean:
    @rm -rf build
    $(GO) clean

.phony: swag
swag:
    $(SWAG) init -g cmd/httpd/main.go --parseInternal -o docs/api/v1/

.phony: fmt
fmt:
    # $(SWAG) fmt $(GOFILES)
    $(GOFMT) -w $(GOFILES)

.phony: test
test:
    $(GO) clean -testcache
    $(GO) mod tidy
    $(GO) test -cover $(GOMODULES)

.phony: update
update:
    $(GO) get -u ./...
    $(GO) mod tidy

.phony: info
info:
    @$(GO) vet $(GOMODULES)
    @$(GO) list $(GOMODULES)
    @$(GO) version
SHELL   := /bin/bash
VERSION := v1.0.0
GOOS    := $(shell go env GOOS)
GOARCH  := $(shell go env GOARCH)

.PHONY: all
all: build

.PHONY: build
build:
	go build -ldflags "-X main.version=$(VERSION)" ./cmd/xjsonl

.PHONY: package
package: clean build
	gzip xjsonl -c > xjsonl_$(VERSION)_$(GOOS)_$(GOARCH).gz
	sha1sum xjsonl_$(VERSION)_$(GOOS)_$(GOARCH).gz > xjsonl_$(VERSION)_$(GOOS)_$(GOARCH).gz.sha1sum

.PHONY: clean
clean:
	rm -f xjsonl

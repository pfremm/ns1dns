PROJ=ns1dns
ORG_PATH=github.com/pfremm
REPO_PATH=$(ORG_PATH)/$(PROJ)
export PATH := $(PWD)/bin:$(PATH)

VERSION ?= $(shell ./scripts/git-version)

$( shell mkdir -p bin )

user=$(shell id -u -n)
group=$(shell id -g -n)

export GOBIN=$(PWD)/bin

LD_FLAGS="-w -X $(REPO_PATH)/version.Version=$(VERSION)"

fmt:
	@./scripts/gofmt $(shell go list ./... | grep -v '/vendor/')

clean:
	@rm -rf bin/

lint:
	@for package in $(shell go list ./... | grep -v '/vendor/' | grep -v '/api' | grep -v '/server/internal'); do \
      golint -set_exit_status $$package $$i || exit 1; \
	done

build: bin/ns1dns

bin/ns1dns: check-go-version
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/ns1dns

release-binary:
	@go build -o $(GOBIN)/ns1dns -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd


check-go-version:
	@./scripts/check-go-version


# Go version to require, run go version to see what version you are using
GO_VERSION := "go{{GoVersion}}"
GO_VERSION ?= $(GO_VERSION)

GO_BIN_DIR?=$(shell dirname `which go`)

.PHONY: check-go-version
check-go-version:
	@if ! go version | grep "$(GO_VERSION)" >/dev/null; then \
        printf "Wrong go version: "; \
        go version; \
        echo "Requires go version: $(GO_VERSION)"; \
        exit 2; \
    fi

.PHONY: build
build: export GOBIN=$(CURDIR)/build/bin
build:
	go install -v $(CURDIR)/cmd/...

.PHONY: generate
generate: export PATH=$(CURDIR)/tools:${GO_BIN_DIR}
generate: tools
	go generate ./...

.PHONY: clean
clean:
	go clean
	rm -rf build/
	rm -rf tmp

.PHONY: lint
lint: install-linter
	golangci-lint --version
	golangci-lint run

.PHONY: install-linter
install-linter:
ifneq ('', '$(shell which golangci-lint)')
	@echo 'golangci-lint is already installed'
else
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.29.0
endif

.PHONY: tools
tools: export GOBIN=$(CURDIR)/tools
tools:
	go install github.com/dmarkham/enumer
	go install github.com/swaggo/swag/cmd/swag

.PHONY: swagger
swagger: tools
	$(CURDIR)/tools/swag init -g internal/rest/router-debug.go -o internal/rest/docs

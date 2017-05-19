
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

GO_PACKAGES = $(shell go list ./... | grep -v vendor | grep -v mocks)
GO_FILES = $(shell find . -name "*.go" | grep -v vendor | uniq)
PROJECT_NAME = $(notdir $(shell pwd))
# The binary to build (just the basename).
BIN := $(PROJECT_NAME)

init:
	@echo "$(OK_COLOR)==> Init$(NO_COLOR)"
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/curtisallen/go-okta
	gometalinter --install

lint:
	@echo "$(OK_COLOR)==> Linting$(NO_COLOR)"
	go list -f '{{.Dir}}' ./... | grep -v 'vendor' | xargs gometalinter --vendored-linters --vendor --concurrency=8

# quick test useful for development
quick-test:
	@echo "$(OK_COLOR)==> Quick Test $(NO_COLOR)"
	go test -cover $(GO_PACKAGES)

# recommended test command
test: lint
	@echo "$(OK_COLOR)==> Testing $(NO_COLOR)"
	go test -race -cover $(GO_PACKAGES)

build:
	@echo "$(OK_COLOR)==> Building $(NO_COLOR)"
	go install

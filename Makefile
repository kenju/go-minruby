NAME := go-minruby
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
		-X 'main.revision=$(REVISION)'

OK_COLOR    = \033[0;32m
ERROR_COLOR = \033[0;31m
WARN_COLOR  = \033[0;33m
NO_COLOR    = \033[m

OK_STRING    = "[OK]"
ERROR_STRING = "[ERROR]"
WARN_STRING  = "[WARNING]"

## Build binaries and run
run: build
	./go-minruby

## Build binaries
build:
	go build -ldflags "$(LDFLAGS)"

## Initial Setup
setup:
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help
	go get github.com/rakyll/gotest


## Run tests
test:
	if LATE_THRESHOLD_MIN=10 gotest ./... -v; then \
		echo "$(OK_COLOR)$(OK_STRING) go test succeeded$(NO_COLOR)"; \
	else \
		echo "$(ERROR_COLOR)$(ERROR_STRING) go test failed$(NO_COLOR)n"; \
	fi


## Show coverage
coverage:
	for pkg in $$(go list ./...); do \
		gotest $$pkg -coverprofile=$(basename $pkg).out ;\
	done
	for pkg in $$(go list ./...); do \
		go tool cover -func=$(basename $pkg).out; \
	done


## Lint source codes
lint:
	go vet ./...
	for d in $$(go list ./...); do \
		golint -set_exit_status ${GOPATH}/src/$$d/*.go || exit $$?; \
	done


## Format source codes
fmt:
	for d in $$(go list ./...); do \
		goimports -w ${GOPATH}/src/$$d; \
	done


## Show help
help:
	@make2help $(MAKEFILE_LIST)


.PHONY: git-hooks setup test coverage lint fmt build help

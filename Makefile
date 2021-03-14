ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BACKEND_DIR := $(ROOT_DIR)/backend

BACKEND_LIBS_PATH := $(BACKEND_DIR)/libs
BACKEND_SOURCE_PATH := $(BACKEND_DIR)/source

tests: backend-tests

check: backend-check

backend-build:
	unset GOPATH && cd $(BACKEND_DIR) && GOMODCACHE=$(BACKEND_LIBS_PATH) go build main.go

backend-run:
	unset GOPATH && cd $(BACKEND_DIR) && GOMODCACHE=$(BACKEND_LIBS_PATH) go run main.go

backend-tests:
	unset GOPATH && cd $(BACKEND_SOURCE_PATH) && GOMODCACHE=$(BACKEND_LIBS_PATH) go test ./... -cover | { grep -v "no test files"; true; }

backend-check:
	unset GOPATH && cd $(BACKEND_SOURCE_PATH) && GOMODCACHE=$(BACKEND_LIBS_PATH) go vet ./...

backend-fmt:
	unset GOPATH && cd $(BACKEND_SOURCE_PATH) && GOMODCACHE=$(BACKEND_LIBS_PATH) go fmt ./...

backend-calc-lines:
	( find $(BACKEND_SOURCE_PATH) -name '*.go' -print0 | xargs -0 cat ) | wc -l

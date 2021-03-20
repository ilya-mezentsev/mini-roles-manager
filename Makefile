ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BACKEND_DIR := $(ROOT_DIR)/backend

DOCKER_COMPOSE_FILE := $(ROOT_DIR)/docker-compose.yaml
PROJECT_NAME := "mini_roles_manager"

BACKEND_LIBS_PATH := $(BACKEND_DIR)/libs
BACKEND_SOURCE_PATH := $(BACKEND_DIR)/source
BACKEND_CONFIG_PATH := $(BACKEND_DIR)/config/main.json

build: backend-build containers-build

run: containers-run

stop: containers-stop

tests: backend-tests

check: backend-check

backend-build:
	unset GOPATH && cd $(BACKEND_DIR) && GOMODCACHE=$(BACKEND_LIBS_PATH) go build main.go

backend-run:
	unset GOPATH && cd $(BACKEND_DIR) && GOMODCACHE=$(BACKEND_LIBS_PATH) go run main.go -config $(BACKEND_CONFIG_PATH)

backend-tests:
	unset GOPATH && cd $(BACKEND_SOURCE_PATH) && GOMODCACHE=$(BACKEND_LIBS_PATH) go test ./... -cover | { grep -v "no test files"; true; }

backend-check:
	unset GOPATH && cd $(BACKEND_SOURCE_PATH) && GOMODCACHE=$(BACKEND_LIBS_PATH) go vet ./...

backend-fmt:
	unset GOPATH && cd $(BACKEND_SOURCE_PATH) && GOMODCACHE=$(BACKEND_LIBS_PATH) go fmt ./...

backend-calc-lines:
	( find $(BACKEND_SOURCE_PATH) -name '*.go' -print0 | xargs -0 cat ) | wc -l

db-run:
	docker-compose -f $(ROOT_DIR)/docker-compose.yaml -p $(PROJECT_NAME) up db

containers-run:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) up

containers-stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) down

containers-build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) build

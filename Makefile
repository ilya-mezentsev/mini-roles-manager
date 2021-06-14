ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BACKEND_DIR := $(ROOT_DIR)/backend
FRONTEND_DIR := $(ROOT_DIR)/frontend

DOCKER_COMPOSE_FILE := $(ROOT_DIR)/docker-compose.yaml
PROJECT_NAME := "mini_roles_manager"

BACKEND_LIBS_PATH := $(BACKEND_DIR)/libs
BACKEND_SOURCE_PATH := $(BACKEND_DIR)/source
BACKEND_CONFIG_PATH := $(BACKEND_DIR)/config/main.json
BACKEND_APP_DATA_PATH := $(BACKEND_DIR)/config/app-data.json

BACKEND_VALIDATION_PKG_PATH := $(BACKEND_SOURCE_PATH)/domains/files/services/validation

FRONTEND_SOURCE_PATH := $(FRONTEND_DIR)/src

build: backend-build frontend-build containers-build

run: containers-run

stop: containers-stop

tests: backend-tests frontend-tests

check: backend-check frontend-check

backend-build:
	unset GOPATH && cd $(BACKEND_DIR) && GOMODCACHE=$(BACKEND_LIBS_PATH) go build -o main

backend-run:
	unset GOPATH && cd $(BACKEND_DIR) && GOMODCACHE=$(BACKEND_LIBS_PATH) ./main -config $(BACKEND_CONFIG_PATH)

backend-run-minimal:
	unset GOPATH && cd $(BACKEND_DIR) && GOMODCACHE=$(BACKEND_LIBS_PATH) ./main -mode minimal -app-data $(BACKEND_APP_DATA_PATH)

backend-help:
	unset GOPATH && cd $(BACKEND_DIR) && GOMODCACHE=$(BACKEND_LIBS_PATH) ./main -help

backend-tests:
	unset GOPATH && cd $(BACKEND_SOURCE_PATH) && GOMODCACHE=$(BACKEND_LIBS_PATH) go test -cover -p 1 ./... | { grep -v "no test files"; true; }

backend-benchmark-file-validation:
	unset GOPATH && cd $(BACKEND_VALIDATION_PKG_PATH) && GOMODCACHE=$(BACKEND_LIBS_PATH) go test -bench . -benchmem -run=^$

backend-check:
	unset GOPATH && cd $(BACKEND_SOURCE_PATH) && GOMODCACHE=$(BACKEND_LIBS_PATH) go vet ./...

backend-fmt:
	unset GOPATH && cd $(BACKEND_SOURCE_PATH) && GOMODCACHE=$(BACKEND_LIBS_PATH) go fmt ./...

backend-calc-lines:
	( find $(BACKEND_SOURCE_PATH) -name '*.go' -print0 | xargs -0 cat ) | wc -l

frontend-build:
	cd $(FRONTEND_DIR) && npm i && npm run build

frontend-build-static:
	cd $(FRONTEND_DIR) && npm run build

frontend-run:
	cd $(FRONTEND_DIR) && npm start

frontend-tests:
	cd $(FRONTEND_DIR) && npm run test -- --watchAll=false --silent

frontend-tests-coverage:
	cd $(FRONTEND_DIR) && npm run test -- \
	--watchAll=false --coverage --silent --collectCoverageFrom=src/**/*.ts \
	--collectCoverageFrom=!src/reportWebVitals.ts \
	--collectCoverageFrom=!src/index.ts \
	--collectCoverageFrom=!src/services/api/shared/request.ts \
	--collectCoverageFrom=!src/services/log/log.ts

frontend-check:
	cd $(FRONTEND_DIR) && npm run lint

frontend-calc-lines:
	( find $(FRONTEND_SOURCE_PATH) -name '*.*' -print0 | xargs -0 cat ) | wc -l

db-run:
	docker-compose -f $(ROOT_DIR)/docker-compose.yaml -p $(PROJECT_NAME) up db

containers-run:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) up

containers-stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) down

containers-build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) build

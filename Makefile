# 定义变量
PWD := $(shell pwd)
SCRIPTS_DIR := $(PWD)/scripts
DEPLOY_DIR := $(PWD)/deployment
BUILD_DIR := $(PWD)/bin
BINARY_NAME := juneblog
GO_FLAGS := -trimpath -ldflags "-s -w"
GIT_SHA := $(shell git rev-parse --short HEAD)

.PHONY: all build clean clean_db build_docker lint test

all: build

build:
	@mkdir -p $(BUILD_DIR)
	@echo "Building $(BINARY_NAME)..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GO_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/main.go
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

lint:
	@echo "Running linter..."
	@golangci-lint run ./...

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned build directory"

clean_db:
	@bash $(SCRIPTS_DIR)/clean_db.sh

build_docker_base:
	@bash $(DEPLOY_DIR)/docker/docker_build.sh $(GIT_SHA) $(USERNAME) $(PASSWORD) $(NAMESPACE)

.DEFAULT_GOAL := all

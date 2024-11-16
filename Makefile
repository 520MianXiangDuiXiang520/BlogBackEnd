# 定义变量
PWD := $(shell pwd)
SCRIPTS_DIR := $(PWD)/scripts
DEPLOY_DIR := $(PWD)/deployment
BUILD_DIR := $(PWD)/bin
BINARY_NAME := juneblog
BUILD_TIME := $(shell date -u '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT := $(shell git rev-parse HEAD)
GO_FLAGS := -trimpath -ldflags "-s -w -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT)"


# 定义 PHONY 目标
.PHONY: all build clean clean_db build_docker

# 默认目标
all: build

# 构建目标
build:
	@mkdir -p $(BUILD_DIR)
	@echo "Building $(BINARY_NAME)..."
	@CGO_ENABLED=0 go build $(GO_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/main.go
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"


# 清理目标
clean:
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned build directory"

# 清理数据库
clean_db:
	@bash $(SCRIPTS_DIR)/clean_db.sh

# 构建 Docker 镜像
build_docker:
	@bash $(DEPLOY_DIR)/docker/docker_build.sh

# 设置默认目标
.DEFAULT_GOAL := all

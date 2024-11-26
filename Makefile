# Variables
APP_NAME := $(shell basename $(CURDIR))
BUILD_DIR := build
DOCKER_IMAGE := $(APP_NAME):latest
CONTAINER_NAME := $(APP_NAME)-container

GO_FILES := $(shell find . -type f -name '*.go')
GO_TESTS := $(shell go list ./...)

# Commands
GO := go
GOLINT := golangci-lint

# Default target
.PHONY: all
all: help

# Show help
.PHONY: help
help: ## Show this help message
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

# Run the application
.PHONY: run
run: ## Run the application
	@echo "Running the application..."
	docker compose up

# Build the application
.PHONY: build
build: ## Compile the application
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) .

# Run tests
.PHONY: test
test: ## Run all tests
	@echo "Running tests..."
	$(GO) test -v $(GO_TESTS)

# Calculate code coverage
.PHONY: coverage
coverage: ## Generate code coverage report
	@echo "Calculating code coverage..."
	$(GO) test -coverprofile=coverage.out $(GO_TESTS)
	@grep -vE "mock|.pb.go" coverage.out > filtered_coverage.out
	$(GO) tool cover -func=filtered_coverage.out

# Run lint
.PHONY: lint
lint: ## Run lint checks
	@echo "Running linter..."
	$(GOLINT) run ./...

# Clean build artifacts
.PHONY: clean
clean: ## Remove build artifacts
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR) ./*.out

# Format code
.PHONY: fmt
fmt: ## Format source code
	@echo "Formatting code..."
	$(GO) fmt ./...

# Install dependencies
.PHONY: deps
deps: ## Install or tidy dependencies
	@echo "Tidying up dependencies..."
	$(GO) mod tidy

# Generate protobuf
.PHONY: gen-pb
gen-pb: ## generate protobuf
	buf generate
	protoc-go-inject-tag -input="./entity/domain/*/*/*.pb.go"

# Generate Swagger specification
.PHONY: gen-swagger
gen-swagger: ## Generate Swagger specification
	@echo "Generating Swagger specification..."
	swag init -g main.go -o ./docs/api

# Docker build
.PHONY: docker-build
docker-build: ## Build the Docker image
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Docker clean
.PHONY: docker-clean
docker-clean: ## Remove Docker image
	@echo "Removing Docker image..."
	docker rmi $(DOCKER_IMAGE)


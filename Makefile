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

# Build the application
.PHONY: build
build: ## Compile the application
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) .

# Run the application
.PHONY: run
run: build ## Run the compiled application
	@echo "Running the application..."
	./$(BUILD_DIR)/$(APP_NAME)

# Run tests
.PHONY: test
test: ## Run all tests
	@echo "Running tests..."
	$(GO) test -v $(GO_TESTS)

# Run lint
.PHONY: lint
lint: ## Run lint checks
	@echo "Running linter..."
	$(GOLINT) run ./...

# Clean build artifacts
.PHONY: clean
clean: ## Remove build artifacts
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

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

# Docker build
.PHONY: docker-build
docker-build: ## Build the Docker image
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Docker run
.PHONY: docker-run
docker-run: ## Run the Docker container
	@echo "Running Docker container..."
	docker run --rm --name $(CONTAINER_NAME) -p 8080:8080 $(DOCKER_IMAGE)

# Docker stop
.PHONY: docker-stop
docker-stop: ## Stop the Docker container
	@echo "Stopping Docker container..."
	docker stop $(CONTAINER_NAME)

# Docker clean
.PHONY: docker-clean
docker-clean: ## Remove Docker image
	@echo "Removing Docker image..."
	docker rmi $(DOCKER_IMAGE)


# Variables
APP_NAME := $(shell basename $(CURDIR))
BUILD_DIR := build
GO_FILES := $(shell find . -type f -name '*.go')
GO_TESTS := $(shell go list ./...)

# Commands
GO := go
GOLINT := golangci-lint

# Default target
.PHONY: all
all: build ## Build the application

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

# Show help
.PHONY: help
help: ## Show this help message
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

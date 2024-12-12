# Variables
APP_NAME=auth0-api-service
GO=go

.PHONY: build run test clean help

# Build the application
build:
	@echo "Building..."
	$(GO) build -o $(APP_NAME) main.go

# Run the application
run:
	@echo "Running..."
	$(GO) run main.go

# Run tests
test:
	@echo "Running tests..."
	$(GO) test ./... -v

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(APP_NAME)
	$(GO) clean

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GO) mod tidy

# Show help
help:
	@echo "Available commands:"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make deps     - Install dependencies" 
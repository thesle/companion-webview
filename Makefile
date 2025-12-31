.PHONY: help install dev build run clean deps

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install Go dependencies
	go mod download
	go mod tidy

dev: ## Run in development mode
	wails dev -tags webkit2_41

build: ## Build the application
	wails build -tags webkit2_41

run: ## Run the built application
	./build/bin/companion-webview

clean: ## Clean build artifacts
	rm -rf build/
	go clean

simple-build: ## Build with standard Go (without Wails CLI)
	go build -tags webkit2_41 -o companion-webview .

simple-run: simple-build ## Build and run with standard Go
	./companion-webview

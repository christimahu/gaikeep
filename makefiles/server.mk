# =====================================================
# SERVER BUILD & RUN COMMANDS
# =====================================================
#
# Purpose: Handles building and running the Gai Keep server
# Usage: Use 'make build' to compile, 'make run' to start the server
# See: docs/workflow.md#common-development-tasks for workflow examples
#
# This makefile contains targets for compiling the server binary,
# running the server, and testing it with demo requests.

# build: Compile the application binary
# Creates the binary in the bin directory with versioning information
.PHONY: build
build:
	@echo "Building the Gai Keep server..."
	@go build -o bin/server cmd/server/main.go

# run: Build and run the server
# Starts the server after ensuring it's compiled with latest changes
.PHONY: run
run: build
	@echo "Running the Gai Keep server..."
	@bin/server

# demo: Send a test entry to the running server
# Usage: make demo "Your test message"
# Note: Requires the server to be already running (use 'make run' in another terminal)
.PHONY: demo
demo:
	@echo "Sending entry to server..."
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		go run scripts/demo.go "Test entry from Makefile"; \
	else \
		go run scripts/demo.go "$(filter-out $@,$(MAKECMDGOALS))"; \
	fi

# This allows arbitrary arguments to be passed to the demo target
%:
	@:

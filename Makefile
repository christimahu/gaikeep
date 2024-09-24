#!/bin/bash

# GAIKeep Makefile - Defines automation tasks for building, testing, and running the GAIKeep application.

# *** Building ***

all: clean test build run  # This target depends on completing clean, test, build, and run steps before running the application.

# Build
# Compiles the Go code and places the app in the bin directory, then makes it executable.
# For more information on the 'go build' command, see the official docs: [https://go.dev/doc/cmd/go/#hdr-Compile_packages_and_dependencies]
# For how the Go compiler works, including building for multiple targets: [https://go.dev/doc/compiler]
build: 
	@echo "Building..."
	@go build -o ./bin/gaikeep . && \
    chmod +x ./bin/gaikeep && \
    echo "Done building!"

# Test
# Runs unit tests for the GAIKeep application
# For more information on the 'go test' command, see the official docs: [https://go.dev/doc/cmd/go/#hdr-Test_packages]
test:
	@echo "Running tests..."
	@go test -v -cover ./internal/... && \
	echo "Done running tests!"

# Clean
# Removes any previously built artifacts
# For more information on the 'rm' command, see the man page: [https://linux.die.net/man/1/rm]
clean:
	@echo "Removing bin folder..."
	@rm -rf bin/ && \
	echo "bin folder removed!"

# Setup
# Prepares the project environment (assuming Go is installed)
# For more information on the 'go mod init' command, see the official docs: [https://go.dev/doc/cmd/go/#hdr-Create_a_new_module]
# For more information on the 'go get' command, see the official docs: [https://go.dev/doc/cmd/go/#hdr-Download_and_install_packages_and_dependencies]
setup:
	@echo "Setting up the GAIKeep project environment..."
	@which go > /dev/null || { echo "Go is not installed. Please install Go first." ; exit 1; }
	@if [ ! -f go.mod ]; then \
        echo "Initializing new Go module..."; \
        go mod init github.com/christimahu/gaikeep; \
    fi
	@echo "Updating project dependencies..."
	@go get -u ./...
	@echo "Installing godotenv package..."
	@go get github.com/joho/godotenv
	# Check if .env file exists and prompt for API key if not
	@if [ ! -f .env ]; then \
        echo "No .env file found." && \
        read -p "Enter your OpenAI API key (Press enter to skip): " apiKey && \
        { [ -z "$$apiKey" ] && echo "No API key entered. Skipping .env file creation." || { \
            echo "Creating .env file with API key..." && \
            echo "OPENAI_API_KEY=$$apiKey" > .env && \
            echo "API key stored in .env file. Please ensure .env is included in your .gitignore to prevent it from being committed."; \
        } }; \
    else \
        echo ".env file already exists. If you need to change the API key, please edit the .env file directly."; \
    fi
	@echo "Setup completed!"

# Run
# Executes the GAIKeep application
run:  
	@echo "Running GaiKeep application..."
	@./bin/gaikeep testgpt


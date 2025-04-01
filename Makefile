# Gai Keep Makefile
#
# Core build and development automation for the Gai Keep project.
# This Makefile follows a modular approach, with specialized functionality
# broken into separate files in the makefiles/ directory.
#
# For detailed explanations of the workflow and philosophy behind this structure,
# see: docs/workflow.md

# =====================================================
# GLOBAL VARIABLES (shared across included makefiles)
# =====================================================

# Go package and directory discovery
# Excludes script directory which contains utility scripts not part of the main application
GO_PACKAGES = $(shell go list ./... | grep -v "/scripts")
GO_DIRS = $(shell find . -type d -not -path "./scripts*" -not -path "*/\\.*" -not -path "./vendor*")

# Go command definitions
# These provide consistent interfaces to go tools across all makefiles
GOCMD       = go
GOBUILD     = $(GOCMD) build
GOTEST      = $(GOCMD) test
GOGET       = $(GOCMD) get
GOMOD       = $(GOCMD) mod
GOFMT       = $(GOCMD) fmt
GOVET       = $(GOCMD) vet
GOLINT      = golangci-lint

# Project metadata
# Basic information that identifies the project and its outputs
PROJECT_NAME = GaiKeep
BINARY_NAME  = server
BINARY_DIR   = bin
MAIN_PKG     = ./cmd/server
PROTO_DIR    = ./proto
# Determine version from git or use "dev" as fallback
VERSION     ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

# Output and artifact directories
# Defines where various build outputs, documentation, and test results will be stored
OUT_DIR         = $(BINARY_DIR)
DOCS_DIR        = web/docs
COVERAGE_DIR    = coverage
COVERAGE_PROFILE= $(COVERAGE_DIR)/coverage.out
COVERAGE_HTML   = $(COVERAGE_DIR)/coverage.html
# Current build time in UTC for versioning
BUILD_TIME      = $(shell date -u '+%Y-%m-%d %H:%M:%S')
# Linker flags for embedding version information
LDFLAGS         = -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# Cross-platform command for opening files
# Falls back to a message if no suitable command is found
OPEN_CMD = open || xdg-open || echo "Open the file manually"

# =====================================================
# LOAD MODULAR MAKEFILE EXTENSIONS
# =====================================================
# Each included makefile handles a specific aspect of the build process.
# See docs/workflow.md for details on the modular architecture.

include makefiles/help.mk      	# Summary of available targets
include makefiles/deps.mk	  	# Install dev dependencies
include makefiles/proto.mk     	# Protobuf code generation
include makefiles/server.mk    	# Compile, run, install binary
include makefiles/test.mk      	# Unit + integration tests and coverage
include makefiles/tools.mk     	# Formatting, vetting, cleanup
include makefiles/docs.mk      	# Documentation generation
include makefiles/lint.mk      	# golangci-lint wrapper

# =====================================================
# DEFAULT TARGET
# =====================================================
# The default target (invoked when running 'make' with no arguments)
# Builds the application for the fastest feedback loop during development
.DEFAULT_GOAL := build

# =====================================================
# AGGREGATE TARGETS
# =====================================================

# all: Run complete build and validation workflow
# This is typically used before committing or for CI pipelines
.PHONY: all
all: clean deps lint test build

# setup: Prepare development environment from scratch
# Useful when first cloning the repo or resetting to a clean state
.PHONY: setup
setup: clean deps proto build
	@echo "Development environment ready."

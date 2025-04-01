# =====================================================
# HELP TARGET
# =====================================================
#
# Purpose: Provides documentation and guidance on available targets
# Usage: Run 'make help' to see a list of available commands
# See: docs/workflow.md#build-system-architecture for more details
#
# This serves as a built-in onboarding tool for new developers.
# The help target lists all major functionality with brief descriptions.

# help: Print a list of common Make targets and descriptions
# This is the default first-stop for new developers to understand the build system
.PHONY: help
help:
	@echo ""
	@echo "Gai Keep Makefile — Available Commands"
	@echo "====================================="
	@echo ""
	@echo "Core Workflow:"
	@echo "  make              - Build only (fastest way to check if it compiles)"
	@echo "  make all          - Full workflow: clean, lint, test, build, docs"
	@echo "  make run          - Compile and launch the app"
	@echo "  make setup        - Prepare dev environment: clean, deps, proto, docs, build"
	@echo ""
	@echo "Testing:"
	@echo "  make unit         - Run unit tests only"
	@echo "  make integration  - Run integration tests (requires Go tag: integration)"
	@echo "  make coverage     - Run tests and open HTML coverage report"
	@echo ""
	@echo "Code Quality:"
	@echo "  make fmt          - Format code using gofmt"
	@echo "  make vet          - Run go vet analysis"
	@echo "  make lint         - Run golangci-lint"
	@echo ""
	@echo "Docs & API:"
	@echo "  make proto        - Generate protobuf + gRPC Go code"
	@echo "  make docs         - Start documentation server at http://localhost:8080"
	@echo ""
	@echo "Housekeeping:"
	@echo "  make clean        - Remove all build artifacts, temp files, test coverage"
	@echo "  make help         - Print this help menu"
	@echo ""

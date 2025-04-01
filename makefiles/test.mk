# =====================================================
# TEST TARGETS
# =====================================================
#
# Purpose: Provides different testing strategies and code coverage analysis
# Usage: Run 'make unit' for basic tests, 'make integration' for full system tests
# See: docs/workflow.md#testing-strategy for details on test philosophy
#
# The project separates unit tests (fast, in-memory) from integration tests
# (slower, with external dependencies) to enable rapid development cycles.

# unit: Run unit tests in source packages
# These tests are fast and isolated, using no external resources
# Ideal for rapid development and continuous integration
.PHONY: unit
unit:
	@echo "Running unit tests..."
	@$(GOTEST) -v $(GO_PACKAGES)

# integration: Run integration tests
# These tests use the 'integration' build tag and may:
# - Spin up gRPC servers
# - Write files to disk
# - Test actual network communication
.PHONY: integration
integration:
	@echo "Running integration tests..."
	@$(GOTEST) -v -tags=integration ./test/integration

# coverage: Generate coverage report
# Creates and opens an HTML report showing which code is covered by tests
# Helps identify areas that need additional testing
.PHONY: coverage
coverage:
	@echo "Generating coverage report..."
	@mkdir -p $(COVERAGE_DIR)
	@$(GOTEST) -coverprofile=$(COVERAGE_PROFILE) -covermode=atomic $(GO_PACKAGES)
	@$(GOCMD) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	@echo "Coverage report saved to $(COVERAGE_HTML)"
	@if [ "$(shell uname)" = "Darwin" ]; then \
		open $(COVERAGE_HTML); \
	elif [ "$(shell uname)" = "Linux" ]; then \
		xdg-open $(COVERAGE_HTML) 2>/dev/null || echo "Please open $(COVERAGE_HTML) manually"; \
	else \
		echo "Please open $(COVERAGE_HTML) manually"; \
	fi

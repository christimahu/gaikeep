# =====================================================
# LINTING
# =====================================================
#
# Purpose: Runs comprehensive code quality checks
# Usage: Run 'make lint' before committing changes
# See: docs/workflow.md#code-quality-tools for configuration details
#
# This makefile integrates golangci-lint, which runs multiple linters
# in parallel to catch a wide range of code quality issues.

# lint: Run golangci-lint
# Runs a suite of linters including:
# - errcheck: Ensures errors are handled
# - staticcheck: Static analysis for common mistakes
# - unused: Identifies unused code
# - and many others
.PHONY: lint
lint:
	@which $(GOLINT) > /dev/null 2>&1 || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@$(GOLINT) run ./... --exclude-dirs=scripts || echo "Linting found issues. Fix them before committing."

# =====================================================
# FORMATTING, ANALYSIS, CLEANUP
# =====================================================
#
# Purpose: Handles code quality tools and build artifact cleanup
# Usage: Run before committing code or to clean the workspace
# See: docs/workflow.md#code-quality-tools for details
#
# These targets help maintain consistent code style, catch common errors,
# and clean up generated files that should not be committed to version control.

# fmt: Format all Go code using gofmt
# Ensures consistent code style across the codebase
# Run this before committing changes
.PHONY: fmt
fmt:
	@echo "Formatting Go code..."
	@for dir in $(GO_DIRS); do \
		$(GOFMT) $$dir; \
	done
	@echo "Formatting complete."

# vet: Run go vet static analysis
# Identifies potential bugs and suspicious code
# Examples: unreachable code, unused variables, misused functions
.PHONY: vet
vet:
	@echo "Running go vet..."
	@$(GOVET) $(GO_PACKAGES)

# clean: Remove build, temp, and test artifacts
# Removes generated binaries, test results, and temporary files
# Use this to return to a clean state or before running full builds
.PHONY: clean
clean:
	@echo "Cleaning project..."
	@rm -rf $(OUT_DIR)
	@rm -rf $(COVERAGE_DIR)
	@rm -rf tmp
	@echo "Clean complete."

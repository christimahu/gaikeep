# =====================================================
# DOCUMENTATION TARGETS
# =====================================================
#
# Purpose: Provides access to Go package documentation
# Usage: Run 'make docs' to start the documentation server
# See: docs/workflow.md#documentation-generation for more on documentation philosophy
#
# This makefile handles the generation and viewing of code documentation
# using Go's built-in documentation system via pkgsite.

# docs: Run pkgsite documentation server locally
# Starts an HTTP server at http://localhost:8080 that displays
# documentation generated from Go code comments
.PHONY: docs
docs:
	@echo "Starting documentation server at http://localhost:8080"
	@echo "Press Ctrl+C to stop"
	@pkgsite -http=:8080

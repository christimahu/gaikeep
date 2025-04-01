# =====================================================
# DEPENDENCIES
# =====================================================

# deps: Install Go module dependencies and tidy
deps:
	@echo "Installing Go dependencies..."
	@$(GOMOD) download
	@$(GOMOD) tidy
	@$(GOCMD) install golang.org/x/pkgsite/cmd/pkgsite@latest
	@echo "Dependencies ready."

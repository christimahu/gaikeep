# Developer Guide

This guide provides information for developers working on the Gai Keep project. It covers development setup, coding standards, and common development tasks.

## Development Environment

### Recommended Tools

- **Editor**: Any text editor with Go support (Neovim, Vim, Emacs, VS Code, GoLand)
- **Terminal**: Any modern terminal with Make support
- **Git Client**: Command line git

### Editor Setup

Gai Keep can be developed with any editor that supports Go. Here are some common setups:

#### Neovim/Vim

For Neovim or Vim users, the following plugins are recommended:
- vim-go or gopls (Go language server)
- treesitter (for improved syntax highlighting)
- telescope (for file navigation)

Example minimal Neovim config for Go development:
```lua
-- Install a plugin manager like packer.nvim first
require('packer').startup(function()
  use 'neovim/nvim-lspconfig'
  use 'nvim-treesitter/nvim-treesitter'
  use 'fatih/vim-go'
end)

-- LSP setup
require('lspconfig').gopls.setup{}

-- Format on save
vim.cmd [[autocmd BufWritePre *.go lua vim.lsp.buf.format()]]
```

#### Other Editors

For other editors, ensure you have:
- Go language server (gopls) integration
- gofmt integration
- Protocol Buffer syntax highlighting

### Project Organization

The project follows standard Go project layout conventions:

- `cmd/` - Application entry points
- `internal/` - Private application code
- `pkg/` - Reusable public packages
- `proto/` - Protocol buffer definitions
- `test/` - Integration tests
- `scripts/` - Utility scripts
- `web/` - Frontend assets
- `docs/` - Documentation
- `makefiles/` - Build system components

## Common Development Tasks

### Setting Up a New Development Environment

```bash
# Clone the repository
git clone https://github.com/ChristiMahu/gaikeep.git
cd gaikeep

# Set up the development environment
make setup
```

### Running Tests

```bash
# Run unit tests
make unit

# Run integration tests
make integration

# Run tests with coverage report
make coverage
```

### Adding a New Feature

1. **Update Protocol Buffer Definition** (if adding API features)
   - Edit `proto/gaikeep/service.proto`
   - Run `make proto` to generate Go code

2. **Implement Server Logic**
   - Add methods to `internal/server/grpc.go`
   - Add tests in `internal/server/grpc_test.go`

3. **Implement Storage Logic** (if needed)
   - Update storage in `internal/storage/file_storage.go`
   - Add tests in `internal/storage/file_storage_test.go`

4. **Test Your Changes**
   - Run `make unit` to run unit tests
   - Run `make integration` to run integration tests

5. **Format and Lint Code**
   - Run `make fmt` to format code
   - Run `make lint` to check for issues

### Debugging

To run the server with more logging:

```bash
# Set debug log level
GAIKEEP_LOG_LEVEL=debug make run
```

To debug with delve:

```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Run debugger
dlv debug cmd/server/main.go
```

## Coding Standards

### Go Style Guide

- Follow standard Go code style (use `make fmt`)
- Use meaningful variable names
- Prefer shorter functions (under 50 lines)
- Add comments for exported functions and packages
- Use error wrapping with context: `fmt.Errorf("failed to do X: %w", err)`

### Commit Message Format

Use descriptive commit messages with a subject line and optional body:

```
area: brief description of change (under 50 chars)

Longer explanation of what this changes and why.
Include any relevant details or context.
Wrap lines at 72 characters.
```

Example areas: `server`, `storage`, `proto`, `docs`, etc.

### Testing Guidelines

- Every package should have tests
- Unit tests should not depend on external resources
- Integration tests should use the `integration` build tag
- Aim for at least 80% code coverage

## Troubleshooting Common Issues

### Protocol Buffer Generation Errors

If you encounter errors with Protocol Buffer generation:

1. Ensure protoc is installed and in your PATH
2. Check that the Go plugins are installed:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
3. Verify your PATH includes $GOPATH/bin
4. Check for syntax errors in your .proto files

### Build Errors

For common build errors:

1. Run `make clean` followed by `make`
2. Update dependencies with `make deps`
3. Check for missing imports
4. Ensure your Go version is 1.22 or higher

### Server Startup Issues

If the server fails to start:

1. Check if another process is using port 50051
2. Ensure you have write permissions to the data directory
3. Check logs for specific error messages

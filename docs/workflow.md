# Gai Keep Development Workflow

This document explains how to use the Makefile system and project layout to build, test, and contribute to Gai Keep. It's designed to complement the shorter descriptions in the Makefile targets by providing broader context and examples.

## Table of Contents

- [Development Philosophy](#development-philosophy)
- [Build System Architecture](#build-system-architecture)
- [Common Development Tasks](#common-development-tasks)
- [Dependency Management](#dependency-management)
- [Testing Strategy](#testing-strategy)
- [Protocol Buffer Workflow](#protocol-buffer-workflow)
- [Documentation Generation](#documentation-generation)
- [Code Quality Tools](#code-quality-tools)
- [Troubleshooting](#troubleshooting)

## Development Philosophy

The Gai Keep build system is designed with these principles in mind:

- **Speed**: Core tasks should be fast to support rapid iteration
- **Modularity**: Functionality is divided into focused makefiles for maintainability
- **Consistency**: Similar patterns are used throughout the codebase
- **Documentation**: Every component should be well-documented

The workflow is built to support both quick development cycles and thorough validation before releases or contributions.

## Build System Architecture

The build system uses a modular approach with:

- A main `Makefile` that defines global variables and includes specialized components
- Individual component makefiles in the `makefiles/` directory
- Consistent naming and organization across all components

This structure keeps the build system maintainable and allows developers to quickly understand where specific functionality lives.

### Component Makefiles

- **help.mk**: Documentation and help target definitions
- **deps.mk**: Dependency management
- **proto.mk**: Protocol buffer code generation
- **server.mk**: Server build and run commands
- **test.mk**: Testing targets for different test types
- **tools.mk**: Code formatting and static analysis
- **docs.mk**: Documentation generation
- **lint.mk**: Code linting with golangci-lint

## Common Development Tasks

### Fast Feedback Loop

For the fastest development cycle:

```bash
# Make a change, then build and run the server
make run
```

### Full Validation

Before submitting changes:

```bash
# Run the complete validation suite
make all
```

### Clean Start

After cloning or when you want a fresh environment:

```bash
# Set up a complete development environment
make setup
```

## Dependency Management

The project uses Go modules for dependency management. The `make deps` target handles several key tasks:

1. Downloads all required dependencies
2. Updates and tidies the go.mod file
3. Installs development tools like pkgsite

### When to Run `make deps`

- After first cloning the repository
- After pulling changes that modify dependencies
- If you're experiencing unexplained build errors

## Testing Strategy

The project uses different types of tests for different purposes:

### Unit Tests

Unit tests are fast, isolated tests that don't require external resources:

```bash
make unit
```

These tests verify individual components in isolation and are the first line of defense against regressions.

### Integration Tests

Integration tests verify that components work together correctly:

```bash
make integration
```

These tests are slower but provide more confidence in the system as a whole. They use the `integration` build tag to keep them separate from unit tests.

### Code Coverage

To see how well the tests cover the codebase:

```bash
make coverage
```

This generates an HTML report showing which lines of code are covered by tests.

## Protocol Buffer Workflow

The project uses Protocol Buffers and gRPC for API definitions:

1. Service definitions are written in `proto/gaikeep/service.proto`
2. Run `make proto` to generate Go code from these definitions
3. The generated code provides both client and server interfaces

Always run `make proto` after changing any `.proto` files.

### Protocol Buffer Workflow with Local Module Resolution

When working with Protocol Buffers in Gai Keep, the build system ensures proper generation and import of protobuf packages through several mechanisms:

#### Local Module Resolution

The project uses a module replacement directive in `go.mod` to ensure that imports of `github.com/ChristiMahu/gaikeep` resolve to the local directory rather than attempting to download from GitHub:

```go
replace github.com/ChristiMahu/gaikeep => ./
```

This directive is automatically added by the `make deps` command if not already present.

#### Protobuf Generation Process

The protobuf generation workflow follows these steps:

1. Define service and message types in `.proto` files in the `proto/gaikeep/` directory
2. Run `make proto` to generate Go code with proper import paths
3. The generated code includes both message types and gRPC service interfaces

The `protoc` command is configured with the correct module path through the `--go_opt=module` and `--go-grpc_opt=module` parameters, ensuring that import statements in the generated code use the correct module path.

#### CI/CD Integration

In the CI/CD pipeline, Protocol Buffer generation occurs before dependency resolution to ensure that all required code is present before the build process begins. The workflow:

1. Adds the replace directive to `go.mod`
2. Installs Protocol Buffer compiler and plugins
3. Generates Go code from `.proto` definitions
4. Proceeds with dependency resolution and testing

This approach ensures consistent behavior between local development and CI/CD environments.

## Documentation Generation

Gai Keep uses Go's standard documentation approach with package and function comments:

```bash
make docs
```

This starts a local documentation server that displays the Go package documentation generated from comments in the code.

### Writing Good Documentation

- Add package-level comments at the top of each file
- Document all exported functions, types, and constants
- Include examples where appropriate
- Keep comments up-to-date when changing code

## Code Quality Tools

Several tools help maintain code quality:

### Formatting

```bash
make fmt
```

Runs `gofmt` to ensure consistent code style across the project.

### Static Analysis

```bash
make vet
```

Runs Go's built-in static analyzer to catch common mistakes.

### Linting

```bash
make lint
```

Runs golangci-lint, which includes multiple linters to catch a wide range of issues.

## Troubleshooting

### Build Errors

If you encounter build errors:

1. Try cleaning and rebuilding: `make clean && make`
2. Update dependencies: `make deps`
3. Check Go version: The project requires Go 1.22+

### Test Failures

If tests are failing:

1. Run the specific failing test with verbose output: `go test -v ./path/to/package -run TestName`
2. Check if you need to regenerate Protocol Buffer code: `make proto`
3. Ensure you have all required dependencies: `make deps`

### Server Won't Start

If the server fails to start:

1. Check for port conflicts (the server uses port 50051 by default)
2. Look for error messages in the output
3. Try running with more verbose logging (set `GAIKEEP_LOG_LEVEL=debug` environment variable)

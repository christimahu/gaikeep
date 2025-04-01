# Gai Keep Documentation

This directory contains documentation for the Gai Keep project.

## Generated Documentation

When you run `make docs-static`, the generated documentation will be placed in the `static/` subdirectory of this folder. These files are not checked into version control and will be removed when running `make clean`.

## Viewing Documentation

To view the most current documentation for Gai Keep:

1. Run the documentation server:
   ```
   make docs
   ```

2. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

## Documentation Structure

The documentation is automatically generated from Go package comments and follows the standard Go documentation conventions:

- Package documentation is taken from the comment block immediately preceding the package declaration
- Function, type, and variable documentation is taken from comments preceding their declarations
- Examples are extracted from `*_test.go` files with the naming pattern `Example*`

## Writing Good Documentation

For guidelines on writing effective Go documentation, please refer to:
- [Godoc: documenting Go code](https://go.dev/blog/godoc)
- [Effective Go: Documentation](https://go.dev/doc/effective_go#commentary)

## Documentation TODOs

- Add developer guide
- Add user guide
- Add API reference

# Testing Strategy

This document outlines the testing approach for the Gai Keep project, including test types, organization, and best practices.

## Testing Philosophy

Gai Keep uses a multi-layered testing approach to ensure code quality:

1. **Unit Tests**: Fast, focused tests for individual components
2. **Integration Tests**: End-to-end tests for system validation
3. **Code Coverage**: Measurement of test thoroughness

The testing strategy aims to balance quick developer feedback with comprehensive validation.

## Test Types

### Unit Tests

Unit tests verify individual components in isolation, mocking external dependencies when necessary. They are:
- Located alongside the code they test (`*_test.go` files)
- Fast (milliseconds to run)
- Independent of each other
- Memory-only (no file I/O or network)

Example:
```go
// internal/model/entry_test.go
func TestEntry_Validate(t *testing.T) {
    entry := Entry{
        ID:        "test-id",
        Timestamp: time.Now().Format(time.RFC3339),
        Content:   "Test content",
    }
    
    err := entry.Validate()
    if err != nil {
        t.Errorf("Validate() error = %v, want nil", err)
    }
}
```

### Integration Tests

Integration tests verify multiple components working together, often testing across package boundaries. They are:
- Located in the `test/integration` directory
- Tagged with `//go:build integration` build constraint
- May use external resources (filesystem, network)
- Test complete workflows

Example:
```go
//go:build integration
// +build integration

// test/integration/grpc_test.go
func TestGRPCStoreEntryIntegration(t *testing.T) {
    // Create a temporary directory
    tmpDir, _ := os.MkdirTemp("", "test-entries")
    defer os.RemoveAll(tmpDir)
    
    // Start a real server, store an entry, verify file exists
    // ...
}
```

## Test Organization

### Directory Structure

```
gaikeep/
├── internal/
│   ├── model/
│   │   ├── entry.go
│   │   └── entry_test.go      // Unit tests
│   ├── storage/
│   │   ├── file_storage.go
│   │   └── file_storage_test.go // Unit tests
│   └── ...
├── test/
│   └── integration/
│       ├── grpc_test.go       // Integration tests
│       └── storage_test.go    // Integration tests
└── ...
```

### Running Tests

```bash
# Run all unit tests
make unit

# Run integration tests
make integration

# Generate coverage report
make coverage
```

## Test Coverage

The project aims for high test coverage:
- Unit tests should cover at least 80% of code
- Integration tests should cover critical paths

Coverage reports can be generated with:
```bash
make coverage
```

This creates an HTML report showing which lines of code are covered by tests.

## Mocking Strategy

Gai Keep uses simple, explicit mocks rather than complex mocking frameworks:

```go
// Example mock storage implementation for testing
type MockStorage struct {
    SaveFunc func(entry Entry) error
    GetFunc  func(id string) (Entry, error)
    ListFunc func() ([]Entry, error)
}

func (m *MockStorage) Save(entry Entry) error {
    return m.SaveFunc(entry)
}

func (m *MockStorage) Get(id string) (Entry, error) {
    return m.GetFunc(id)
}

func (m *MockStorage) List() ([]Entry, error) {
    return m.ListFunc()
}
```

## Test Data Management

For test data:
- Use meaningful example data (not foo/bar)
- Generate temporary files/directories when needed
- Clean up after tests finish
- Don't commit test-generated files

Example:
```go
func TestStorage(t *testing.T) {
    // Create temporary test directory
    tmpDir, err := os.MkdirTemp("", "gaikeep-test")
    if err != nil {
        t.Fatalf("Failed to create temp dir: %v", err)
    }
    
    // Always clean up after the test
    defer os.RemoveAll(tmpDir)
    
    // Test implementation...
}
```

## Continuous Integration

In the future, the project will implement CI pipelines to:
- Run all tests on every pull request
- Enforce minimum code coverage
- Run linters and other code quality tools

## Testing Best Practices

1. **Write tests first** when possible
2. **One assertion per test** for clear failure messages
3. **Use table-driven tests** for multiple similar test cases
4. **Include both positive and negative tests**
5. **Keep tests fast** to maintain fast feedback loops
6. **Cleanup test resources** to avoid interference between tests

Example of table-driven tests:
```go
func TestEntry_Validate(t *testing.T) {
    tests := []struct {
        name    string
        entry   Entry
        wantErr bool
    }{
        {
            name: "valid entry",
            entry: Entry{
                ID:        "test-id",
                Timestamp: time.Now().Format(time.RFC3339),
                Content:   "Test content",
            },
            wantErr: false,
        },
        {
            name: "missing ID",
            entry: Entry{
                ID:        "",
                Timestamp: time.Now().Format(time.RFC3339),
                Content:   "Test content",
            },
            wantErr: true,
        },
        // More test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.entry.Validate()
            if (err != nil) != tt.wantErr {
                t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

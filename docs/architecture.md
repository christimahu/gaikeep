# Gai Keep Architecture

This document describes the architecture of the Gai Keep application, providing an overview of the system's components and how they interact.

## System Overview

Gai Keep follows a modular, layered architecture that separates concerns and promotes maintainability. The application is designed around a gRPC server that processes client requests, with file-based storage for persistence.

## Component Diagram

```
┌───────────────┐      ┌───────────────┐      ┌───────────────┐
│  gRPC Client  │ ──→  │  gRPC Server  │ ──→  │ File Storage  │
└───────────────┘      └───────────────┘      └───────────────┘
                              │
                              ↓
                       ┌───────────────┐
                       │   Data Model  │
                       └───────────────┘
```

## Key Components

### gRPC Server (`internal/server`)

The gRPC server implements the protocol buffer service defined in `proto/gaikeep/service.proto`. It handles incoming requests and coordinates between the application logic and storage layer.

Key responsibilities:
- Processing client requests
- Input validation
- Coordinating data flow
- Error handling

Key file: `internal/server/grpc.go`

### Storage Layer (`internal/storage`)

The storage layer handles persistence of entries to the filesystem as JSON files. It abstracts the details of file handling from the rest of the application.

Key responsibilities:
- Saving entries to disk
- Retrieving entries by ID
- Listing all entries
- File organization and naming

Key file: `internal/storage/file_storage.go`

### Data Model (`internal/model`)

The data model defines the core data structures used throughout the application. It includes the Entry struct and validation logic.

Key responsibilities:
- Defining data structures
- Data validation
- JSON serialization/deserialization

Key file: `internal/model/entry.go`

### Application Logic (`internal/app`)

The application layer ties everything together, creating and configuring the server and its dependencies.

Key responsibilities:
- Server lifecycle management
- Configuration
- Dependency injection

Key file: `internal/app/server.go`

## Data Flow

1. Client sends a gRPC request (SaveEntry, GetEntry, or ListEntries)
2. Server receives the request and validates inputs
3. Server passes data to the storage layer
4. Storage layer performs file operations
5. Result returns back through the layers to the client

## Storage Format

Entries are stored as individual JSON files with the following naming pattern:
```
entry_<UUID>_<TIMESTAMP>.json
```

Each file contains a JSON object with the Entry structure, including:
- ID (UUID)
- Timestamp (RFC3339 format)
- Content (text)
- Tags (optional array of strings)

## Configuration

The application uses a simple configuration mechanism (`internal/config`) that supports:
- Environment variable overrides (GAIKEEP_DATA_DIR)
- Default values when not specified

## Error Handling

Error handling is implemented with Go's standard error pattern:
- Errors are propagated up through the call stack
- Each layer adds context to errors
- The server converts internal errors to appropriate gRPC status codes

## Logging

A simple logging package (`pkg/logger`) provides consistent logging throughout the application with different severity levels:
- Info for operational events
- Error for recoverable errors
- Fatal for non-recoverable errors

## Future Expansion Points

The architecture is designed to allow for future enhancements:
- Additional storage backends could be implemented
- Authentication could be added to the gRPC server
- The data model could be extended with additional fields or types
- A frontend could connect to the gRPC API

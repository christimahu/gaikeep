// Package main provides the entry point for the Gai Keep server.
//
// This file initializes the application server, configures it with
// appropriate middleware and storage handlers, and starts the gRPC
// service. It also handles graceful shutdown and error reporting.
// The server listens on the default port 50051 and implements the
// interface defined in the service.proto file.
package main

import (
	"github.com/ChristiMahu/gaikeep/internal/app"
	"github.com/ChristiMahu/gaikeep/pkg/logger"
)

// main initializes and starts the Gai Keep server.
// It creates a new server instance on the default port (:50051),
// then blocks until the server stops or an error occurs.
// If server creation or startup fails, the application will log
// a fatal error and exit. The server handles gRPC requests for
// saving, retrieving, and listing AI-generated content entries.
func main() {
	// Create a new server instance
	server, err := app.NewServer(":50051")
	if err != nil {
		logger.LogFatal("failed to create server: " + err.Error())
	}

	// Start the server (this blocks until server is stopped)
	if err := server.Start(); err != nil {
		logger.LogFatal("server error: " + err.Error())
	}
}

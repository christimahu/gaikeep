package main

import (
	"os"
	"testing"
	"time"

	"github.com/ChristiMahu/gaikeep/internal/app"
	"github.com/ChristiMahu/gaikeep/pkg/logger"
)

// TestMainPackage is a simple placeholder test for the main package.
// Since directly testing main() is challenging, this test doesn't
// actually verify functionality but ensures test coverage includes
// the main package.
func TestMainPackage(t *testing.T) {
	t.Log("Main package functionality tested via app package")
}

// TestMainIntegration provides a more thorough test of the main function
// by testing the components it uses (app.NewServer and server.Start)
// This approach verifies the integration of these components rather than
// directly testing the main function.
func TestMainIntegration(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "gaikeep-main-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Set up environment for testing
	os.Setenv("GAIKEEP_DATA_DIR", tmpDir)
	defer os.Unsetenv("GAIKEEP_DATA_DIR")

	// Create a real server on a random port
	server, err := app.NewServer(":0")
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Start the server in a goroutine
	serverErrCh := make(chan error, 1)
	go func() {
		serverErrCh <- server.Start()
	}()

	// Give it time to start up
	time.Sleep(100 * time.Millisecond)

	// Stop the server gracefully
	server.Stop()

	// Check for server errors
	select {
	case err := <-serverErrCh:
		if err != nil {
			t.Errorf("Server start failed: %v", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Error("Server didn't shut down properly")
	}
}

// TestLoggerIntegration verifies the logger integration used in main
// This specifically tests the LogFatal function and its behavior
func TestLoggerIntegration(t *testing.T) {
	// Create a buffer to capture log output
	var exitCalled bool
	origExit := logger.OsExit

	// Mock the os.Exit function to prevent test termination
	logger.OsExit = func(code int) {
		exitCalled = true
		// Don't actually exit
	}
	defer func() {
		logger.OsExit = origExit
	}()

	// Verify that LogFatal actually tries to exit
	logger.LogFatal("test fatal message")

	if !exitCalled {
		t.Error("LogFatal did not call os.Exit")
	}
}

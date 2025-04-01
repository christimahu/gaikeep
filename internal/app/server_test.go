package app

import (
	"os"
	"testing"
	"time"
)

// TestNewServer contains test cases for the NewServer function.
// It verifies that the server can be properly initialized with
// valid configuration and that appropriate errors are returned
// for invalid configurations.
func TestNewServer(t *testing.T) {
	// Test valid configuration with a temporary directory
	t.Run("valid config", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "gaikeep-server-test")
		if err != nil {
			t.Fatalf("failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		// Set temp dir as data directory
		os.Setenv("GAIKEEP_DATA_DIR", tmpDir)
		defer os.Unsetenv("GAIKEEP_DATA_DIR")

		server, err := NewServer(":50051")
		if err != nil {
			t.Fatalf("NewServer() error = %v", err)
		}

		if server == nil {
			t.Fatal("NewServer() returned nil")
		}

		if server.address != ":50051" {
			t.Errorf("NewServer() address = %v, want %v", server.address, ":50051")
		}

		if server.config == nil {
			t.Error("NewServer() config is nil")
		}

		if server.storage == nil {
			t.Error("NewServer() storage is nil")
		}

		if server.grpc == nil {
			t.Error("NewServer() grpc is nil")
		}
	})

	// Test with invalid configuration to ensure error handling
	t.Run("invalid config", func(t *testing.T) {
		// Set invalid directory to trigger error
		os.Setenv("GAIKEEP_DATA_DIR", "/invalid-path-that-cannot-exist/with-no-permissions")
		defer os.Unsetenv("GAIKEEP_DATA_DIR")

		_, err := NewServer(":50051")
		if err == nil {
			t.Error("NewServer() with invalid config expected error, got nil")
		}
	})
}

// TestServerStart validates the server's startup and shutdown functionality.
// It tests that the server can start successfully and be gracefully stopped.
// This test starts the actual server in a goroutine and uses channels to
// coordinate shutdown and error detection.
func TestServerStart(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "gaikeep-server-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Set temp dir as data directory
	os.Setenv("GAIKEEP_DATA_DIR", tmpDir)
	defer os.Unsetenv("GAIKEEP_DATA_DIR")

	// Use a random high port to avoid conflicts
	server, err := NewServer(":0")
	if err != nil {
		t.Fatalf("NewServer() error = %v", err)
	}

	// Start server in background
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Start()
	}()

	// Give it a moment to start up
	time.Sleep(100 * time.Millisecond)

	// Stop the server
	server.Stop()

	// Check for any errors
	select {
	case err := <-errCh:
		if err != nil {
			t.Errorf("server.Start() error = %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Error("server.Start() did not complete after stop")
	}
}

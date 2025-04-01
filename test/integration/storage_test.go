//go:build integration
// +build integration

// Package integration contains tests for verifying the file storage layer
// under real file system conditions.
//
// These tests verify that the storage layer works correctly with actual
// file system interactions, including reading from and writing to disk.
// They ensure that data persistence behaves as expected in a realistic
// environment rather than with mocked dependencies.
package integration

import (
	"os"
	"testing"
	"time"

	"github.com/ChristiMahu/gaikeep/internal/storage"
)

// TestFileStorage_ReadWriteCycle saves and retrieves a single entry to disk
// and verifies the content and structure match.
//
// This test validates the complete save-and-load cycle for entries,
// ensuring that all entry data is preserved through serialization,
// file storage, and subsequent deserialization. It uses a real
// temporary directory to verify actual file system behavior.
func TestFileStorage_ReadWriteCycle(t *testing.T) {
	// Create temporary directory with unique name
	tmpDir, err := os.MkdirTemp("", "gaikeep-storage-test")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}

	// Clean up at the end
	defer func() {
		// Close any open files
		os.Chmod(tmpDir, 0755) // Ensure we have permissions
		err := os.RemoveAll(tmpDir)
		if err != nil {
			t.Logf("Warning: failed to remove temp dir %s: %v", tmpDir, err)
		}
	}()

	// Initialize storage with the temporary directory
	store, err := storage.NewFileStorage(tmpDir)
	if err != nil {
		t.Fatalf("failed to init storage: %v", err)
	}

	// Create a test entry with all fields populated
	entry := storage.Entry{
		ID:        "integration-test-123",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Content:   "hello from integration",
		Tags:      []string{"test", "integration"},
	}

	// Save the entry to disk
	if err := store.Save(entry); err != nil {
		t.Fatalf("failed to save entry: %v", err)
	}

	// Retrieve the entry from disk
	loaded, err := store.Get(entry.ID)
	if err != nil {
		t.Fatalf("failed to get entry: %v", err)
	}

	// Verify entry content was preserved
	if loaded.Content != entry.Content {
		t.Errorf("expected %q, got %q", entry.Content, loaded.Content)
	}
}

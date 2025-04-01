package storage

import (
	"os"
	"path/filepath"
	"testing"
)

// TestNewFileStorage verifies that the FileStorage constructor properly
// initializes storage instances with the provided directory and handles
// error conditions appropriately.
func TestNewFileStorage(t *testing.T) {
	t.Run("with valid directory", func(t *testing.T) {
		// Create a temporary directory for testing
		tmpDir, err := os.MkdirTemp("", "storage-test")
		if err != nil {
			t.Fatalf("failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		// Initialize storage with the temporary directory
		fs, err := NewFileStorage(tmpDir)
		if err != nil {
			t.Fatalf("NewFileStorage() error = %v", err)
		}
		if fs.Directory != tmpDir {
			t.Errorf("NewFileStorage() directory = %v, want %v", fs.Directory, tmpDir)
		}
	})

	t.Run("with empty directory", func(t *testing.T) {
		// Test that an empty directory path is rejected
		_, err := NewFileStorage("")
		if err == nil {
			t.Errorf("NewFileStorage() with empty directory expected error, got nil")
		}
	})
}

// TestFileStorage_Save verifies that entries can be successfully saved to disk
// and that validation errors are properly handled.
func TestFileStorage_Save(t *testing.T) {
	// Create testing directory structure
	tmpDir := "tmp/test-storage"
	_ = os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	fs, err := NewFileStorage(tmpDir)
	if err != nil {
		t.Fatalf("failed to create file storage: %v", err)
	}

	t.Run("valid entry", func(t *testing.T) {
		// Create a valid entry for testing
		entry := Entry{
			ID:        "entry123",
			Timestamp: "2025-03-30T00:00:00Z",
			Content:   "Test content",
			Tags:      []string{"test"},
		}

		// Save the entry
		err = fs.Save(entry)
		if err != nil {
			t.Fatalf("failed to save entry: %v", err)
		}

		// Check file exists
		path := filepath.Join(tmpDir, "entry_entry123_2025-03-30T00:00:00Z.json")
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Save() did not create file at %s", path)
		}
	})

	t.Run("invalid entry", func(t *testing.T) {
		// Create an invalid entry (empty ID) for testing
		entry := Entry{
			ID:        "", // Invalid: empty ID
			Timestamp: "2025-03-30T00:00:00Z",
			Content:   "Test content",
		}

		// Attempt to save the invalid entry
		err = fs.Save(entry)
		if err == nil {
			t.Errorf("Save() with invalid entry expected error, got nil")
		}
	})
}

// TestFileStorage_Get verifies that entries can be successfully retrieved by ID
// and that appropriate errors are returned when entries cannot be found.
func TestFileStorage_Get(t *testing.T) {
	// Create testing directory structure
	tmpDir := "tmp/test-storage"
	_ = os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	fs, err := NewFileStorage(tmpDir)
	if err != nil {
		t.Fatalf("failed to create file storage: %v", err)
	}

	// Create and save a test entry
	entry := Entry{
		ID:        "entry123",
		Timestamp: "2025-03-30T00:00:00Z",
		Content:   "Test content",
		Tags:      []string{"test"},
	}

	err = fs.Save(entry)
	if err != nil {
		t.Fatalf("failed to save entry: %v", err)
	}

	t.Run("existing entry", func(t *testing.T) {
		// Attempt to load the entry by ID
		loadedEntry, err := fs.Get("entry123")
		if err != nil {
			t.Fatalf("failed to get entry: %v", err)
		}

		// Verify content matches the original
		if loadedEntry.Content != entry.Content {
			t.Errorf("expected content %q, got %q", entry.Content, loadedEntry.Content)
		}
	})

	t.Run("non-existent entry", func(t *testing.T) {
		// Attempt to load an entry with an ID that doesn't exist
		_, err := fs.Get("nonexistent")
		if err == nil {
			t.Errorf("Get() with non-existent ID expected error, got nil")
		}
	})
}

// TestFileStorage_List verifies that all entries in the storage directory
// can be listed successfully and returned as a collection.
func TestFileStorage_List(t *testing.T) {
	// Create testing directory structure
	tmpDir := "tmp/test-storage"
	_ = os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	fs, err := NewFileStorage(tmpDir)
	if err != nil {
		t.Fatalf("failed to create file storage: %v", err)
	}

	// Save a couple of entries
	entry1 := Entry{
		ID:        "entry1",
		Timestamp: "2025-03-30T00:00:00Z",
		Content:   "Test content 1",
	}
	entry2 := Entry{
		ID:        "entry2",
		Timestamp: "2025-03-30T00:00:00Z",
		Content:   "Test content 2",
	}

	if err := fs.Save(entry1); err != nil {
		t.Fatalf("failed to save entry1: %v", err)
	}
	if err := fs.Save(entry2); err != nil {
		t.Fatalf("failed to save entry2: %v", err)
	}

	entries, err := fs.List()
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if len(entries) != 2 {
		t.Errorf("List() returned %v entries, want 2", len(entries))
	}

	// Check both entries are in the list
	found1, found2 := false, false
	for _, entry := range entries {
		if entry.ID == "entry1" {
			found1 = true
		}
		if entry.ID == "entry2" {
			found2 = true
		}
	}

	if !found1 || !found2 {
		t.Errorf("List() missing entries, found1=%v, found2=%v", found1, found2)
	}
}

// TestEntry_Validate verifies that Entry validation correctly identifies
// valid and invalid entries based on their field values.
func TestEntry_Validate(t *testing.T) {
	t.Run("valid entry", func(t *testing.T) {
		entry := Entry{
			ID:        "test",
			Timestamp: "2025-03-30T00:00:00Z",
			Content:   "content",
		}
		if err := entry.Validate(); err != nil {
			t.Errorf("Validate() error = %v", err)
		}
	})

	t.Run("invalid timestamp", func(t *testing.T) {
		entry := Entry{
			ID:        "test",
			Timestamp: "invalid",
			Content:   "content",
		}
		if err := entry.Validate(); err == nil {
			t.Errorf("Validate() with invalid timestamp expected error, got nil")
		}
	})
}

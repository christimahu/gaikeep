package server

import (
	"context"
	"os"
	"testing"

	"github.com/ChristiMahu/gaikeep/internal/storage"
	pb "github.com/ChristiMahu/gaikeep/proto/gaikeep"
)

// TestNew verifies that the New function correctly initializes
// a GaiKeepServer with the provided storage implementation.
// It ensures that all dependencies are properly set up and
// the server is ready to handle requests.
func TestNew(t *testing.T) {
	// Create real storage for testing
	tmpDir, err := os.MkdirTemp("", "server-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	store, err := storage.NewFileStorage(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	s := New(store)
	if s == nil {
		t.Fatal("New() returned nil")
	}

	if s.store != store {
		t.Errorf("New() did not set store correctly")
	}
}

// TestServerMethods tests the gRPC service methods of the GaiKeepServer.
// It verifies that the SaveEntry, GetEntry, and ListEntries methods
// function correctly with valid inputs and return appropriate errors
// with invalid inputs. This test uses a real file storage backend
// with a temporary directory to ensure end-to-end functionality.
func TestServerMethods(t *testing.T) {
	// Set up real storage for testing
	tmpDir, err := os.MkdirTemp("", "server-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	store, err := storage.NewFileStorage(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	server := New(store)
	ctx := context.Background()

	// Test SaveEntry
	content := "Test content"
	saveResp, err := server.SaveEntry(ctx, &pb.SaveEntryRequest{
		Content: content,
	})
	if err != nil {
		t.Fatalf("SaveEntry() error = %v", err)
	}
	if saveResp.Id == "" {
		t.Errorf("SaveEntry() returned empty ID")
	}

	// Test GetEntry
	getResp, err := server.GetEntry(ctx, &pb.GetEntryRequest{
		Id: saveResp.Id,
	})
	if err != nil {
		t.Fatalf("GetEntry() error = %v", err)
	}
	if getResp.Content != content {
		t.Errorf("GetEntry() content = %v, want %v", getResp.Content, content)
	}

	// Test ListEntries
	listResp, err := server.ListEntries(ctx, &pb.ListEntriesRequest{})
	if err != nil {
		t.Fatalf("ListEntries() error = %v", err)
	}
	if len(listResp.Entries) != 1 {
		t.Errorf("ListEntries() returned %v entries, want 1", len(listResp.Entries))
	}

	// Test GetEntry with invalid ID
	_, err = server.GetEntry(ctx, &pb.GetEntryRequest{
		Id: "invalid-id",
	})
	if err == nil {
		t.Errorf("GetEntry() with invalid ID expected error, got nil")
	}
}

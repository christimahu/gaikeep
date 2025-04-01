//go:build integration
// +build integration

// Package integration contains full end-to-end integration tests
// that start real gRPC servers and test file persistence.
//
// These tests verify the complete system functionality by testing
// multiple components working together in realistic scenarios.
// Unlike unit tests, these tests use actual resources like
// filesystem storage and network connections.
package integration

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ChristiMahu/gaikeep/internal/server"
	"github.com/ChristiMahu/gaikeep/internal/storage"
	pb "github.com/ChristiMahu/gaikeep/proto/gaikeep"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TestGRPCStoreEntryIntegration starts a gRPC server and verifies
// that storing an entry creates a corresponding JSON file on disk.
//
// This test verifies the complete flow from gRPC API call to file
// creation, ensuring that all components work together correctly.
// It uses a temporary directory for storage to avoid interference
// with other tests or existing data.
func TestGRPCStoreEntryIntegration(t *testing.T) {
	// Create a temporary directory that won't conflict with other tests
	tmpDir, err := os.MkdirTemp("", "gaikeep-test-entries")
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

	// Initialize the storage system with the temporary directory
	store, err := storage.NewFileStorage(tmpDir)
	if err != nil {
		t.Fatalf("failed to create file storage: %v", err)
	}

	// Create a network listener on an available port
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	// Initialize the gRPC server with our storage
	grpcServer := grpc.NewServer()
	pb.RegisterGaiKeepServer(grpcServer, server.New(store))

	// Start the server in a goroutine
	go func() {
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	// Ensure server is stopped at the end of the test
	defer func() {
		grpcServer.GracefulStop()
		lis.Close()
	}()

	// Create a client connection to our server
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}

	// Ensure connection is closed
	defer conn.Close()

	// Create a client for our API
	client := pb.NewGaiKeepClient(conn)

	// Set a timeout for API calls
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test content to save
	testContent := "integration test payload"

	// Call the SaveEntry RPC to store the content
	resp, err := client.SaveEntry(ctx, &pb.SaveEntryRequest{
		Content: testContent,
	})
	if err != nil {
		t.Fatalf("SaveEntry RPC failed: %v", err)
	}

	// Verify that a file was created for the entry
	matches, err := filepath.Glob(filepath.Join(tmpDir, fmt.Sprintf("entry_%s_*.json", resp.Id)))
	if err != nil || len(matches) == 0 {
		t.Errorf("expected file for entry %s not found", resp.Id)
	}
}

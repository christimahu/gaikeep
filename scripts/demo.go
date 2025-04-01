// Package main provides a demonstration client for the Gai Keep API.
//
// This script connects to a running Gai Keep server and demonstrates basic
// functionality by sending an entry to be stored. It serves as both a demonstration
// of how to use the Gai Keep API from Go code and as a quick test to verify that
// the server is operational and accessible.
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ChristiMahu/gaikeep/proto/gaikeep"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// main is the entry point for the demo client script.
// It connects to the Gai Keep server using gRPC, sends a content entry
// to be saved, and logs the resulting entry ID.
//
// The script accepts an optional command-line argument that specifies
// the content text to save. If no argument is provided, it uses a default
// test message.
func main() {
	// Default content
	content := "Test entry from Makefile"

	// If an argument is provided, use it as the content
	if len(os.Args) > 1 {
		content = os.Args[1]
	}

	// Connect to the gRPC server using insecure credentials
	// Note: In a production environment, secure credentials should be used
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client using the connection
	client := gaikeep.NewGaiKeepClient(conn)

	// Set a timeout for the API call
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Call SaveEntry RPC to store the content
	resp, err := client.SaveEntry(ctx, &gaikeep.SaveEntryRequest{
		Content: content,
	})
	if err != nil {
		log.Fatalf("Failed to save entry: %v", err)
	}

	// Log the successful result
	log.Printf("Successfully saved entry with ID: %s", resp.Id)
}

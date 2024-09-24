package storage

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSaveResponse(t *testing.T) {
	outputDir := "./output"
	// Attempt to remove and recreate the output directory for clean state testing
	os.RemoveAll(outputDir)
	defer os.RemoveAll(outputDir) // Clean up after test

	response := "Test response content"
	err := SaveResponse(response)
	if err != nil {
		t.Fatalf("Failed to save response: %v", err)
	}

	// Verify the file was created with the expected content
	timestamp := time.Now().Format("20060102-150405")
	expectedFile := fmt.Sprintf("%s/response-%s.txt", outputDir, timestamp)

	content, err := os.ReadFile(expectedFile)
	if err != nil {
		t.Fatalf("Failed to read the file: %v", err)
	}
	if string(content) != response {
		t.Errorf("Expected file content %q; got %q", response, string(content))
	}
}

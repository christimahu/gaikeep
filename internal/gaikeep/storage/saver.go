package storage

import (
	"fmt"
	"os"
	"time"
)

// SaveResponse saves a string response to a file in the output directory.
// It ensures that the directory exists and handles file creation and writing.
func SaveResponse(response string) error {
	dir := "./output" // Path to the output directory from the project root
	// Check if the directory exists, create it if not
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	// Create a file with a timestamped filename
	timestamp := time.Now().Format("20060102-150405")
	filename := fmt.Sprintf("%s/response-%s.txt", dir, timestamp)
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Write the response to the file
	_, err = file.WriteString(response)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}

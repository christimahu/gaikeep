// Package storage provides the file-based persistence layer for entries.
//
// This package implements the storage interface for Gai Keep, handling the
// reading and writing of entry data to the filesystem. It organizes entries
// as individual JSON files with consistent naming conventions and provides
// operations for saving, retrieving, and listing entries.
package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ChristiMahu/gaikeep/pkg/logger"
)

// FileStorage implements the storage.Storage interface for file-based storage.
// It manages the persistence of entries to JSON files in a specified directory,
// handling all file operations and maintaining a consistent structure for
// the stored data.
type FileStorage struct {
	Directory string
}

// NewFileStorage creates a new FileStorage instance with the given directory.
// It ensures the specified directory exists, creating it if necessary.
//
// The directory parameter specifies the path where entry files will be stored.
// If the directory does not exist, it will be created with permissions 0755.
//
// Returns an initialized FileStorage and any error encountered during creation.
// Common errors include permission issues or invalid directory paths.
func NewFileStorage(directory string) (*FileStorage, error) {
	if directory == "" {
		return nil, errors.New("directory cannot be empty")
	}

	// Create directory if it doesn't exist
	if err := os.MkdirAll(directory, 0755); err != nil {
		logger.LogError("Failed to create storage directory: " + err.Error())
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	logger.LogInfo("File storage initialized at: " + directory)

	return &FileStorage{
		Directory: directory,
	}, nil
}

// Save persists an entry to a file in the configured storage directory.
// The entry is validated before saving and is written as a formatted JSON file.
//
// The filename follows the pattern: entry_<ID>_<Timestamp>.json
//
// Returns an error if validation fails, the file cannot be created, or
// JSON encoding encounters an error.
func (fs *FileStorage) Save(entry Entry) error {
	if err := entry.Validate(); err != nil {
		logger.LogError("Invalid entry: " + err.Error())
		return fmt.Errorf("invalid entry: %w", err)
	}

	fileName := fmt.Sprintf("entry_%s_%s.json", entry.ID, entry.Timestamp)
	filePath := filepath.Join(fs.Directory, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		logger.LogError("Failed to create file: " + err.Error())
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(entry); err != nil {
		logger.LogError("Failed to encode entry: " + err.Error())
		return fmt.Errorf("failed to encode entry: %w", err)
	}

	logger.LogInfo("Entry saved to file: " + filePath)
	return nil
}

// Get retrieves an entry by its ID from the file storage.
// It searches for files matching the pattern entry_<ID>_*.json
// and decodes the first matching file into an Entry structure.
//
// Returns the found Entry and nil if successful, or an empty Entry
// and an error if the entry is not found or cannot be decoded.
func (fs *FileStorage) Get(id string) (Entry, error) {
	logger.LogInfo("Looking for entry with ID: " + id)

	files, err := filepath.Glob(filepath.Join(fs.Directory, fmt.Sprintf("entry_%s_*.json", id)))
	if err != nil || len(files) == 0 {
		logger.LogError("Entry not found: " + id)
		return Entry{}, fmt.Errorf("entry not found: %s", id)
	}

	filePath := files[0]
	file, err := os.Open(filePath)
	if err != nil {
		logger.LogError("Failed to open file: " + err.Error())
		return Entry{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var entry Entry
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&entry); err != nil {
		logger.LogError("Failed to decode entry: " + err.Error())
		return Entry{}, fmt.Errorf("failed to decode entry: %w", err)
	}

	logger.LogInfo("Entry retrieved successfully")
	return entry, nil
}

// List returns all entries from the file storage.
// It reads all files matching the pattern entry_*.json in the
// storage directory and decodes them into Entry structures.
//
// Returns a slice of all entries found and nil if successful,
// or nil and an error if files cannot be read or decoded.
func (fs *FileStorage) List() ([]Entry, error) {
	logger.LogInfo("Listing all entries in directory: " + fs.Directory)

	var entries []Entry
	files, err := filepath.Glob(filepath.Join(fs.Directory, "entry_*.json"))
	if err != nil {
		logger.LogError("Failed to list files: " + err.Error())
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	logger.LogInfo(fmt.Sprintf("Found %d entry files", len(files)))

	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			logger.LogError("Failed to open file: " + err.Error())
			return nil, fmt.Errorf("failed to open file: %w", err)
		}

		var entry Entry
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&entry); err != nil {
			file.Close()
			logger.LogError("Failed to decode entry: " + err.Error())
			return nil, fmt.Errorf("failed to decode entry: %w", err)
		}

		file.Close()
		entries = append(entries, entry)
	}

	return entries, nil
}

// Entry represents a single data object stored in the system.
// This is a local version of the model.Entry struct that allows
// the storage package to operate independently of the model package.
type Entry struct {
	ID        string   `json:"id"`
	Timestamp string   `json:"timestamp"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags,omitempty"`
}

// Validate checks that the entry has valid, non-empty fields.
// It ensures that:
// - The ID is not empty
// - The Content is not empty
// - The Timestamp is not empty and is in RFC3339 format
//
// Returns an error if any validation check fails.
func (e *Entry) Validate() error {
	if e.ID == "" {
		return errors.New("entry ID cannot be empty")
	}
	if e.Content == "" {
		return errors.New("entry content cannot be empty")
	}
	if e.Timestamp == "" {
		return errors.New("entry timestamp cannot be empty")
	}
	_, err := time.Parse(time.RFC3339, e.Timestamp)
	if err != nil {
		return fmt.Errorf("invalid timestamp: %w", err)
	}
	return nil
}

// Package config provides application configuration loading and management.
//
// This package handles loading configuration settings from environment variables
// and providing default values when needed. It centralizes all configuration
// logic to make it easier to modify and extend the application's configuration
// options without changing code throughout the codebase.
package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config holds the runtime configuration for Gai Keep.
// It contains all settings that can be customized through environment
// variables or command-line flags.
type Config struct {
	// Directory is the base path where entries are persisted as JSON.
	// This can be overridden with the GAIKEEP_DATA_DIR environment variable.
	Directory string
}

// Load initializes a Config with default values.
// It checks for a GAIKEEP_DATA_DIR environment variable override.
// If the environment variable is not set, it uses a "data" directory
// within the current working directory.
//
// Returns a populated Config struct and any error encountered during loading.
// Common errors include inability to determine the working directory.
func Load() (*Config, error) {
	dir := os.Getenv("GAIKEEP_DATA_DIR")
	if dir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("could not get working directory: %w", err)
		}
		dir = filepath.Join(cwd, "data")
	}

	return &Config{
		Directory: dir,
	}, nil
}

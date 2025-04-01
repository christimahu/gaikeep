package config

import (
	"os"
	"path/filepath"
	"testing"
)

// TestLoad verifies the configuration loading functionality.
// These tests cover both default configuration and environment variable
// overrides to ensure the Config object is properly initialized.
func TestLoad(t *testing.T) {
	// Test default directory configuration when no environment variable is set
	t.Run("default directory", func(t *testing.T) {
		// Clear any existing environment variable to test default behavior
		os.Unsetenv("GAIKEEP_DATA_DIR")

		cfg, err := Load()
		if err != nil {
			t.Fatalf("Load() error = %v", err)
		}

		// Verify default data directory is correctly set relative to current directory
		cwd, _ := os.Getwd()
		expected := filepath.Join(cwd, "data")
		if cfg.Directory != expected {
			t.Errorf("Load() Directory = %v, want %v", cfg.Directory, expected)
		}
	})

	// Test directory override using the environment variable
	t.Run("environment variable override", func(t *testing.T) {
		// Set test directory path in environment variable
		testDir := "/tmp/test-gaikeep-data"
		os.Setenv("GAIKEEP_DATA_DIR", testDir)
		defer os.Unsetenv("GAIKEEP_DATA_DIR")

		cfg, err := Load()
		if err != nil {
			t.Fatalf("Load() error = %v", err)
		}

		// Verify environment variable value is respected
		if cfg.Directory != testDir {
			t.Errorf("Load() Directory = %v, want %v", cfg.Directory, testDir)
		}
	})
}

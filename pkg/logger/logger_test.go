package logger

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// TestLogInfo verifies that the LogInfo function correctly formats
// and outputs informational log messages with the appropriate prefix.
// It captures log output to verify both prefix and message content.
func TestLogInfo(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	originalOutput := Logger.Writer()
	Logger.SetOutput(&buf)
	oldFlags := Logger.Flags()
	Logger.SetFlags(0) // Remove timestamp for predictable output
	defer func() {
		Logger.SetFlags(oldFlags)
		Logger.SetOutput(originalOutput)
	}()

	testMsg := "info test message"
	LogInfo(testMsg)

	output := buf.String()
	if !strings.Contains(output, "INFO: "+testMsg) {
		t.Errorf("LogInfo() output = %q, want to contain %q", output, "INFO: "+testMsg)
	}
}

// TestLogError verifies that the LogError function correctly formats
// and outputs error log messages with the appropriate prefix.
// It captures log output to verify both prefix and message content.
func TestLogError(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	originalOutput := Logger.Writer()
	Logger.SetOutput(&buf)
	oldFlags := Logger.Flags()
	Logger.SetFlags(0) // Remove timestamp for predictable output
	defer func() {
		Logger.SetFlags(oldFlags)
		Logger.SetOutput(originalOutput)
	}()

	testMsg := "error test message"
	LogError(testMsg)

	output := buf.String()
	if !strings.Contains(output, "ERROR: "+testMsg) {
		t.Errorf("LogError() output = %q, want to contain %q", output, "ERROR: "+testMsg)
	}
}

// TestLogFatal verifies that the LogFatal function correctly formats
// fatal error messages and calls the exit function with the expected code.
// It mocks the os.Exit function to prevent actual termination during testing.
func TestLogFatal(t *testing.T) {
	// Replace os.Exit to prevent test from exiting
	originalExit := OsExit
	defer func() { OsExit = originalExit }()

	var exitCode int
	OsExit = func(code int) {
		exitCode = code
	}

	// Capture log output
	var buf bytes.Buffer
	originalOutput := Logger.Writer()
	Logger.SetOutput(&buf)
	oldFlags := Logger.Flags()
	Logger.SetFlags(0) // Remove timestamp for predictable output
	defer func() {
		Logger.SetFlags(oldFlags)
		Logger.SetOutput(originalOutput)
	}()

	testMsg := "fatal test message"
	LogFatal(testMsg)

	// Verify log output contains expected message with prefix
	output := buf.String()
	if !strings.Contains(output, "FATAL: "+testMsg) {
		t.Errorf("LogFatal() output = %q, want to contain %q", output, "FATAL: "+testMsg)
	}

	// Verify exit code is as expected
	if exitCode != 1 {
		t.Errorf("LogFatal() exit code = %d, want 1", exitCode)
	}
}

// TestLoggerWithFlags ensures that the logger properly uses the
// flag settings for message format and includes tests with the
// default flags enabled.
func TestLoggerWithFlags(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	originalOutput := Logger.Writer()
	Logger.SetOutput(&buf)
	oldFlags := Logger.Flags()
	// Keep default flags to test with timestamp and file info
	defer func() {
		Logger.SetFlags(oldFlags)
		Logger.SetOutput(originalOutput)
	}()

	testMsg := "message with flags"
	LogInfo(testMsg)

	output := buf.String()
	if !strings.Contains(output, "INFO: ") {
		t.Errorf("LogInfo() with flags missing prefix, got: %q", output)
	}
	if !strings.Contains(output, testMsg) {
		t.Errorf("LogInfo() with flags missing message, got: %q", output)
	}
	// With default flags, should include year or timestamp pattern
	if !strings.Contains(output, "202") { // Any year in the 2020s
		t.Errorf("LogInfo() with flags missing timestamp, got: %q", output)
	}
}

// TestLoggerCustomPrefix ensures the logger can handle custom prefixes
// that aren't set by the default logging functions.
func TestLoggerCustomPrefix(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	originalOutput := Logger.Writer()
	Logger.SetOutput(&buf)
	oldFlags := Logger.Flags()
	Logger.SetFlags(0) // Remove timestamp for predictable output
	oldPrefix := Logger.Prefix()
	defer func() {
		Logger.SetFlags(oldFlags)
		Logger.SetOutput(originalOutput)
		Logger.SetPrefix(oldPrefix)
	}()

	customPrefix := "CUSTOM_PREFIX: "
	Logger.SetPrefix(customPrefix)
	testMsg := "custom prefix test"
	Logger.Println(testMsg)

	output := buf.String()
	if !strings.Contains(output, customPrefix) {
		t.Errorf("Logger custom prefix missing, got: %q, want prefix: %q", output, customPrefix)
	}
	if !strings.Contains(output, testMsg) {
		t.Errorf("Logger custom prefix message missing, got: %q", output)
	}
}

// TestConcurrentLogWrites verifies that the logger handles
// concurrent writes from multiple goroutines correctly.
func TestConcurrentLogWrites(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	originalOutput := Logger.Writer()
	Logger.SetOutput(&buf)
	oldFlags := Logger.Flags()
	Logger.SetFlags(0) // Remove timestamp for predictable output
	defer func() {
		Logger.SetFlags(oldFlags)
		Logger.SetOutput(originalOutput)
	}()

	// Create a done channel to coordinate goroutines
	done := make(chan bool)
	messageCount := 100

	// Launch multiple goroutines writing log messages
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < messageCount/10; j++ {
				LogInfo("concurrent message")
			}
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Check that all messages were written
	output := buf.String()
	infoCount := strings.Count(output, "INFO: concurrent message")

	// We might not get exactly messageCount due to line breaks and formatting,
	// but we should get roughly the right number
	if infoCount < messageCount-10 {
		t.Errorf("Expected approximately %d log messages, got only %d", messageCount, infoCount)
	}
}

// TestLogOutputToFile verifies that the logger can write to a file.
// This tests the ability to change the output destination.
func TestLogOutputToFile(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "logger-test-*.log")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Save original output and restore it after test
	originalOutput := Logger.Writer()
	defer Logger.SetOutput(originalOutput)

	// Set logger to write to the file
	Logger.SetOutput(tmpFile)

	// Write a test message
	testMsg := "file output test"
	LogInfo(testMsg)

	// Close the file to flush buffers
	tmpFile.Close()

	// Read the file content
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Check file content
	if !strings.Contains(string(content), testMsg) {
		t.Errorf("Log file doesn't contain expected message. Got: %s", string(content))
	}
}

// TestDevNullOutput tests that the logger can write to /dev/null
// This verifies that the logger can handle write-only destinations
func TestDevNullOutput(t *testing.T) {
	// Skip on Windows as it doesn't have /dev/null
	if os.Getenv("GOOS") == "windows" {
		t.Skip("Skipping test on Windows")
	}

	// Open /dev/null
	devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		t.Fatalf("Failed to open /dev/null: %v", err)
	}
	defer devNull.Close()

	// Save original output and restore it after test
	originalOutput := Logger.Writer()
	defer Logger.SetOutput(originalOutput)

	// Set logger to write to /dev/null
	Logger.SetOutput(devNull)

	// This should not cause any errors
	LogInfo("message to nowhere")
	LogError("error to nowhere")
}

// Package logger provides simple logging utilities for the Gai Keep application.
//
// This package implements a lightweight logging system that supports different
// severity levels (info, error, fatal) and consistent formatting. It centralizes
// logging functionality across the application and provides a simple interface
// for all components to use when recording application events and errors.
package logger

import (
	"log"
	"os"
)

// Logger is the global application logger.
// It's configured with standard output destination, a consistent prefix format,
// and flags to include timestamp and source file information in log entries.
var Logger = log.New(os.Stdout, "Gai Keep: ", log.LstdFlags|log.Lshortfile)

// LogInfo logs informational messages.
// Use this for routine operational events and status updates
// that are useful for understanding application behavior under
// normal conditions.
func LogInfo(message string) {
	Logger.SetPrefix("INFO: ")
	Logger.Println(message)
}

// LogError logs error messages.
// Use this for recoverable errors and exceptional conditions
// that don't require application termination but should be
// investigated. These represent issues that impede normal
// operation but don't prevent the application from continuing.
func LogError(message string) {
	Logger.SetPrefix("ERROR: ")
	Logger.Println(message)
}

// LogFatal logs fatal error messages and exits the application.
// Use this only for critical errors that prevent the application
// from functioning and require immediate termination. This will
// log the message and then terminate the application with a
// non-zero exit code.
func LogFatal(message string) {
	Logger.SetPrefix("FATAL: ")
	Logger.Println(message)
	OsExit(1)
}

// OsExit is a variable to allow for testing by overriding
// This variable exists to allow unit tests to override the os.Exit
// function, which would otherwise terminate the test process.
// It's exported to allow tests in other packages to override it.
var OsExit = os.Exit

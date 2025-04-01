// Package model defines the core data model used for entries in the system.
//
// This package contains the fundamental data structures that represent
// entries stored by Gai Keep. It includes validation logic, serialization
// methods, and utility functions for working with these data structures.
// The model package is used throughout the application whenever entry
// data needs to be manipulated.
package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Entry represents a single data object stored in the system.
// An Entry contains the content to be saved along with metadata
// such as an identifier, creation timestamp, and optional tags
// for categorization and searching.
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

// MarshalJSON implements custom JSON encoding for Entry.
// This method allows for customized JSON serialization of
// Entry objects, following the standard Marshaler interface
// from the encoding/json package.
//
// Returns the JSON-encoded bytes and any encoding error.
func (e Entry) MarshalJSON() ([]byte, error) {
	type Alias Entry
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(e),
	})
}

// UnmarshalJSON implements custom JSON decoding for Entry.
// This method allows for customized JSON deserialization of
// Entry objects, following the standard Unmarshaler interface
// from the encoding/json package.
//
// It accepts JSON-encoded bytes and returns any decoding error.
func (e *Entry) UnmarshalJSON(data []byte) error {
	type Alias Entry
	aux := &struct {
		Alias
	}{
		Alias: (Alias)(*e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*e = Entry(aux.Alias)
	return nil
}

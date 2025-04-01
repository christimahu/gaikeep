package model

import (
	"encoding/json"
	"testing"
	"time"
)

// TestEntry_Validate verifies the validation logic for Entry objects.
// It tests both valid entries and various invalid configurations to ensure
// that the validation rules are properly enforced.
func TestEntry_Validate(t *testing.T) {
	validTime := time.Now().UTC().Format(time.RFC3339)

	// Define test cases covering both valid and invalid configurations
	tests := []struct {
		name    string
		entry   Entry
		wantErr bool
	}{
		{
			name: "valid entry",
			entry: Entry{
				ID:        "test-id",
				Timestamp: validTime,
				Content:   "Test content",
				Tags:      []string{"tag1", "tag2"},
			},
			wantErr: false,
		},
		{
			name: "missing ID",
			entry: Entry{
				ID:        "",
				Timestamp: validTime,
				Content:   "Test content",
			},
			wantErr: true,
		},
		{
			name: "missing timestamp",
			entry: Entry{
				ID:        "test-id",
				Timestamp: "",
				Content:   "Test content",
			},
			wantErr: true,
		},
		{
			name: "invalid timestamp",
			entry: Entry{
				ID:        "test-id",
				Timestamp: "not-a-timestamp",
				Content:   "Test content",
			},
			wantErr: true,
		},
		{
			name: "missing content",
			entry: Entry{
				ID:        "test-id",
				Timestamp: validTime,
				Content:   "",
			},
			wantErr: true,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.entry.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestEntry_MarshalJSON verifies that Entry objects can be properly
// serialized to JSON format. It ensures that all fields are correctly
// included in the JSON output and can be unmarshaled back to an Entry.
func TestEntry_MarshalJSON(t *testing.T) {
	entry := Entry{
		ID:        "test-id",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Content:   "Test content",
		Tags:      []string{"tag1", "tag2"},
	}

	// Marshal the entry to JSON
	data, err := entry.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON() error = %v", err)
	}

	// Validate the JSON can be unmarshaled back
	var unmarshaled Entry
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify all fields were correctly preserved
	if unmarshaled.ID != entry.ID {
		t.Errorf("MarshalJSON() ID = %v, want %v", unmarshaled.ID, entry.ID)
	}
	if unmarshaled.Timestamp != entry.Timestamp {
		t.Errorf("MarshalJSON() Timestamp = %v, want %v", unmarshaled.Timestamp, entry.Timestamp)
	}
	if unmarshaled.Content != entry.Content {
		t.Errorf("MarshalJSON() Content = %v, want %v", unmarshaled.Content, entry.Content)
	}
	if len(unmarshaled.Tags) != len(entry.Tags) {
		t.Errorf("MarshalJSON() Tags length = %v, want %v", len(unmarshaled.Tags), len(entry.Tags))
	}
}

// TestEntry_UnmarshalJSON verifies that JSON data can be properly
// deserialized into Entry objects. It ensures that the unmarshaling
// process correctly populates all fields of the Entry struct.
func TestEntry_UnmarshalJSON(t *testing.T) {
	// Create JSON test data with all fields
	jsonData := `{
		"id": "test-id",
		"timestamp": "2025-03-30T12:34:56Z",
		"content": "Test content",
		"tags": ["tag1", "tag2"]
	}`

	// Unmarshal the JSON to an Entry
	var entry Entry
	err := entry.UnmarshalJSON([]byte(jsonData))
	if err != nil {
		t.Fatalf("UnmarshalJSON() error = %v", err)
	}

	// Verify all fields were correctly populated
	if entry.ID != "test-id" {
		t.Errorf("UnmarshalJSON() ID = %v, want %v", entry.ID, "test-id")
	}
	if entry.Timestamp != "2025-03-30T12:34:56Z" {
		t.Errorf("UnmarshalJSON() Timestamp = %v, want %v", entry.Timestamp, "2025-03-30T12:34:56Z")
	}
	if entry.Content != "Test content" {
		t.Errorf("UnmarshalJSON() Content = %v, want %v", entry.Content, "Test content")
	}
	if len(entry.Tags) != 2 || entry.Tags[0] != "tag1" || entry.Tags[1] != "tag2" {
		t.Errorf("UnmarshalJSON() Tags = %v, want %v", entry.Tags, []string{"tag1", "tag2"})
	}
}

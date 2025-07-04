package database

import (
	"testing"
)

func TestConnect(t *testing.T) {
	// Test with invalid connection string
	_, err := Connect("invalid-connection-string")
	if err == nil {
		t.Error("Expected error for invalid connection string")
	}
	
	// Note: We don't test with a real database connection in unit tests
	// Integration tests would handle that separately
}
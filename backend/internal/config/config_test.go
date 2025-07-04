package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Test default values
	cfg := Load()
	
	if cfg.DatabaseURL == "" {
		t.Error("DatabaseURL should have a default value")
	}
	
	if cfg.JWTSecret == "" {
		t.Error("JWTSecret should have a default value")
	}
	
	if cfg.Environment == "" {
		t.Error("Environment should have a default value")
	}
}

func TestLoadWithEnvVars(t *testing.T) {
	// Set environment variables
	os.Setenv("DB_URL", "test-db-url")
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("ENV", "test")
	
	defer func() {
		os.Unsetenv("DB_URL")
		os.Unsetenv("JWT_SECRET")
		os.Unsetenv("ENV")
	}()
	
	cfg := Load()
	
	if cfg.DatabaseURL != "test-db-url" {
		t.Errorf("Expected DatabaseURL 'test-db-url', got '%s'", cfg.DatabaseURL)
	}
	
	if cfg.JWTSecret != "test-secret" {
		t.Errorf("Expected JWTSecret 'test-secret', got '%s'", cfg.JWTSecret)
	}
	
	if cfg.Environment != "test" {
		t.Errorf("Expected Environment 'test', got '%s'", cfg.Environment)
	}
}

func TestGetEnv(t *testing.T) {
	// Test with existing env var
	os.Setenv("TEST_VAR", "test-value")
	defer os.Unsetenv("TEST_VAR")
	
	result := getEnv("TEST_VAR", "default")
	if result != "test-value" {
		t.Errorf("Expected 'test-value', got '%s'", result)
	}
	
	// Test with non-existing env var
	result = getEnv("NON_EXISTING_VAR", "default")
	if result != "default" {
		t.Errorf("Expected 'default', got '%s'", result)
	}
}
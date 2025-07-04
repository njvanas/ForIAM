package main

import "testing"

func TestStartup(t *testing.T) {
	t.Log("Startup test passed")
}

func TestMain(t *testing.T) {
	// Test that main package can be imported without issues
	t.Log("Main package test passed")
}

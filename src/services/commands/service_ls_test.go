package services

import (
	"bytes"
	"os"
	"testing"
)

func TestLsCommand(t *testing.T) {
	// Store the original working directory to restore later
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get original working directory: %v", err)
	}

	// Prepare a temporary directory for testing
	dir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(dir) // Cleanup after test

	// Change the working directory to the temp directory for testing
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("Failed to change directory to temp dir: %v", err)
	}
	defer os.Chdir(originalDir) // Ensure the original directory is restored after the test

	// Create some test files in the temporary directory
	filesToCreate := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, filename := range filesToCreate {
		if err := os.WriteFile(filename, []byte("test content"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	tests := []struct {
		args           []string
		expectedOutput string
		expectedError  string
	}{
		{args: []string{}, expectedOutput: "file1.txt file2.txt file3.txt\n", expectedError: ""},
		{args: []string{"something"}, expectedOutput: "", expectedError: "Ls does not accept arguments\n"},
	}

	for _, tt := range tests {
		var outputWriter, errorWriter bytes.Buffer
		LsCommand(tt.args, &outputWriter, &errorWriter)

		if outputWriter.String() != tt.expectedOutput {
			t.Errorf("For args %v, expected output %q, got %q", tt.args, tt.expectedOutput, outputWriter.String())
		}

		if errorWriter.String() != tt.expectedError {
			t.Errorf("For args %v, expected error %q, got %q", tt.args, tt.expectedError, errorWriter.String())
		}
	}
}

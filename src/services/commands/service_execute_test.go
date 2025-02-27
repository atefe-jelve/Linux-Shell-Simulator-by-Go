package services

import (
	"bytes"
	"testing"
)

func TestStatusCommand(t *testing.T) {
	tests := []struct {
		command string
		args           []string
		expectedOutput string
		expectedError  string
	}{
		{command: "", args: []string{}, expectedOutput: "", expectedError: "No command provided\n"},
		{command: "echo",args: []string{"hello"}, expectedOutput: "hello\n", expectedError: ""},
	}

	for _, tt := range tests {
		var outputWriter, errorWriter bytes.Buffer
		ExecuteCommand(tt.command, tt.args, &outputWriter, &errorWriter)
		if outputWriter.String() != tt.expectedOutput {
			t.Errorf("Expected status %q, got %q", tt.expectedOutput, outputWriter )
		}
		if errorWriter.String() != tt.expectedError {
			t.Errorf("Expected error %q, got %q", tt.expectedError, errorWriter.String())
		}
	}
}

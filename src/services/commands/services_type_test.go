package services

import (
	"bytes"
	"testing"
)

func TestTypeCommand(t *testing.T) {
	tests := []struct {
		args           []string
		expectedOutput string
		expectedError  string
	}{
		{args: []string{}, expectedOutput: "", expectedError: "Please insert a command to check\n"},
		{args: []string{"fakecommand"}, expectedOutput: "", expectedError: "fakecommand: command not found\n"},
		{args: []string{"exit"}, expectedOutput: "exit is a shell builtin\n", expectedError: ""},
		{args: []string{"invalid"}, expectedOutput: "git is /usr/bin/git\n", expectedError: ""},
	}

	for _, tt := range tests {
		var outputWriter, errorWriter bytes.Buffer
		TypeCommand(tt.args, &outputWriter, &errorWriter)
		if outputWriter.String() != tt.expectedOutput {
			t.Errorf("Expected output %q, got %q", tt.expectedOutput, outputWriter)
		}
		if errorWriter.String() != tt.expectedError {
			t.Errorf("Expected error %q, got %q", tt.expectedError, errorWriter.String())
		}
	}
}

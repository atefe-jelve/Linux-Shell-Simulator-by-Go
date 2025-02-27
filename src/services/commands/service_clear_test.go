package services

import (
	"bytes"
	"testing"
)

func TestClearCommand(t *testing.T) {

	tests := []struct {
		args           []string
		expectedOutput string
		expectedError  string
	}{
		{args: []string{}, expectedOutput: "\x1b[H\x1b[2J", expectedError: ""},
		{args: []string{"something"}, expectedOutput: "", expectedError: "ClearCommand does not accept any arguments\n"},
	}

	for _, tt := range tests {
		var outputWriter, errorWriter bytes.Buffer
		ClearCommand(tt.args, &outputWriter, &errorWriter)

		if outputWriter.String() != tt.expectedOutput {
			t.Errorf("For args %v, expected output %q, got %q", tt.args, tt.expectedOutput, outputWriter.String())
		}

		if errorWriter.String() != tt.expectedError {
			t.Errorf("For args %v, expected error %q, got %q", tt.args, tt.expectedError, errorWriter.String())
		}
	}
}

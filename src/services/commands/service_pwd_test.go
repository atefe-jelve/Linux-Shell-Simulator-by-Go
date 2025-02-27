package services

import (
	"bytes"
	"os"
	"testing"
)

func TestPwdCommand(t *testing.T) {
	expectedDir, _ := os.Getwd()

	tests := []struct {
		args           []string
		expectedOutput string
		expectedError  string
	}{
		{args: []string{}, expectedOutput: expectedDir + "\n", expectedError: ""},
		{args: []string{"something"}, expectedOutput: "", expectedError: "pwd does not accept arguments\n"},
	}

	for _, tt := range tests {
		var outputWriter, errorWriter bytes.Buffer
		PwdCommand(tt.args, &outputWriter, &errorWriter)

		if outputWriter.String() != tt.expectedOutput {
			t.Errorf("For args %v, expected output %q, got %q", tt.args, tt.expectedOutput, outputWriter.String())
		}

		if errorWriter.String() != tt.expectedError {
			t.Errorf("For args %v, expected error %q, got %q", tt.args, tt.expectedError, errorWriter.String())
		}
	}
}

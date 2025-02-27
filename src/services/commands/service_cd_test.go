package services

import (
	"bytes"
	"testing"
)

func TestCdCommands(t *testing.T) {
	tests := []struct {
		args           []string
		expectedOutput string
		expectedError  string
	}{
		{args: []string{}, expectedOutput: "", expectedError: "No directory provided\n"},
		{args: []string{"/home", "/home/Desktop"}, expectedOutput: "", expectedError: "Only one directory can be provided\n"},
		{args: []string{"/home"}, expectedOutput: "Changed directory to: /home\n", expectedError: ""},
		{args: []string{"AA"}, expectedOutput: "", expectedError: "Error changing directory to AA: chdir AA: no such file or directory\n"},
	}

	for _, tt := range tests {
		var outputWriter, errorWriter bytes.Buffer
		CdCommand(tt.args, &outputWriter, &errorWriter)

		if outputWriter.String() != tt.expectedOutput {
			t.Errorf("For args %v, expected output %q, got %q", tt.args, tt.expectedOutput, outputWriter.String())
		}

		if errorWriter.String() != tt.expectedError {
			t.Errorf("For args %v, expected error %q, got %q", tt.args, tt.expectedError, errorWriter.String())
		}
	}
}

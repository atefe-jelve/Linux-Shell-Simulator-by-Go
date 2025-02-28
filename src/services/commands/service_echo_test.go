package services

import (
	"bytes"
	"os"
	"testing"
)

func TestEchoCommands(t *testing.T) {

	tests := []struct {
		args           []string
		expectedOutput string
		expectedError  string
	}{
		{args: []string{"hello"}, expectedOutput: "hello\n", expectedError: ""},
		{args: []string{"no$PATH"}, expectedOutput: "no" + os.Getenv("PATH") + "\n", expectedError: ""},
		{args: []string{"no yes"}, expectedOutput: "no yes\n", expectedError: ""},
		{args: []string{"'no$PATH'"}, expectedOutput: "no$PATH\n", expectedError: ""},
		{args: []string{"\"This is a \"quoted\" word.\""}, expectedOutput: "This is a \"quoted\" word.\n", expectedError: ""},
		{args: []string{"\"ab \\a \\$ \\` \\\" \\\\\""}, expectedOutput: "ab \\a $ \\` \" \n", expectedError: ""},
	}

	for _, tt := range tests {
		var outputWriter, errorWriter bytes.Buffer
		EchoCommand(tt.args, &outputWriter, &errorWriter)

		if outputWriter.String() != tt.expectedOutput {
			t.Errorf("For args %v, expected output %q, got %q", tt.args, tt.expectedOutput, outputWriter.String())
		}

		if errorWriter.String() != tt.expectedError {
			t.Errorf("For args %v, expected error %q, got %q", tt.args, tt.expectedError, errorWriter.String())
		}
	}
}

package services

import (
	"bytes"
	"projectshell/src/databases"
	"testing"
)

func TestAddUser(t *testing.T) {

	databases.SetupTestDB(&User{})
	databases.SetDB(databases.GetTestDB()) // Override the global `db`
	defer func() {}()
	
	tests := []struct {
		args           []string
		expectedOutput string
		expectedError  string
	}{
		{args: []string{"user_test1", "123"}, expectedOutput: "User created successfully.\n", expectedError: ""},
		{args: []string{"user_test2"}, expectedOutput: "User created successfully.\n", expectedError: ""},
		{args: []string{"user_test2"}, expectedOutput: "", expectedError: "Duplicate user exists with this username.\n"},
		{args: []string{""}, expectedOutput: "Not enough arguments provided.\n", expectedError: ""},
	}

	for _, tt := range tests {
		var outputWriter, errorWriter bytes.Buffer
		AddUserCommand(tt.args, &outputWriter, &errorWriter)

		if outputWriter.String() != tt.expectedOutput {
			t.Errorf("For args %v, expected output %q, got %q", tt.args, tt.expectedOutput, outputWriter.String())
		}

		if errorWriter.String() != tt.expectedError {
			t.Errorf("For args %v, expected error %q, got %q", tt.args, tt.expectedError, errorWriter.String())
		}
	}
}

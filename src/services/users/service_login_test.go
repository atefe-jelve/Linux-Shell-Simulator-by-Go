package services

import (
	"bytes"
	"projectshell/src/databases"
	"testing"
)

func setupTestDatabase() {
	databases.SetupTestDB(&User{})
	databases.SetDB(databases.GetTestDB())

	db := databases.GetDB()

	// Insert a test user into the in-memory test DB.
	db.Create(&User{
		UserName: "testuser",
		Password: "password123",
	})
}

func clearSessionCache() {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	sessionCache = make(map[string]string) 
}

func TestLoginCommand(t *testing.T) {
	setupTestDatabase()
	clearSessionCache()

	tests := []struct {
		args           []string
		expectedOutput string
		expectedError  string
		expectedUser   string // after successful login, check current user
	}{
		{args: []string{}, expectedOutput: "Not enough arguments provided.\n", expectedError: "", expectedUser: ""},
		{args: []string{"clean"}, expectedOutput: "Not enough arguments provided.\n", expectedError: "", expectedUser: ""},
		{args: []string{"wronguser", "password123"}, expectedOutput: "Invalid credentials\n", expectedError: "", expectedUser: ""},
		{args: []string{"testuser", "wrongpassword"}, expectedOutput: "Invalid credentials\n", expectedError: "", expectedUser: ""},
		{args: []string{"testuser", "password123"}, expectedOutput: "", expectedError: "", expectedUser: "testuser"},
	}

	for _, tt := range tests {
		var outputWriter, errorWriter bytes.Buffer

		LoginCommand(tt.args, &outputWriter, &errorWriter)

		if outputWriter.String() != tt.expectedOutput {
			t.Errorf("For args %v, expected output %q, got %q", tt.args, tt.expectedOutput, outputWriter.String())
		}

		if errorWriter.String() != tt.expectedError {
			t.Errorf("For args %v, expected error %q, got %q", tt.args, tt.expectedError, errorWriter.String())
		}

		currentUser := GetCurrentUser()
		if currentUser != tt.expectedUser {
			t.Errorf("For args %v, expected current user to be %q, got %q", tt.args, tt.expectedUser, currentUser)
		}

		clearSessionCache()
	}
}

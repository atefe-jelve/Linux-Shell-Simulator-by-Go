package services

import (
	"bytes"
	"testing"
)

func TestLogoutCommand(t *testing.T) {
	// Simulate logged-in user
	AddUserSession("testuser")
	SetCurrentUser("testuser")

	// Confirm user is logged in 
	if !IsSessionValid("testuser") || GetCurrentUser() != "testuser" {
		t.Fatalf("user not logged in correctly")
	}

	//  logout 
	var outputWriter, errorWriter bytes.Buffer
	LogoutCommand(nil, &outputWriter, &errorWriter)

	// Check user is logged out
	if IsSessionValid("testuser") {
		t.Errorf("User session should be invalid after logout")
	}

	if GetCurrentUser() != "" {
		t.Errorf("Current user should be empty after logout, got %q", GetCurrentUser())
	}

	// expectedOutput := "Logged out successfully.\n"
	// if outputWriter.String() != expectedOutput {
	//     t.Errorf("Expected output %q, got %q", expectedOutput, outputWriter.String())
	// }
}

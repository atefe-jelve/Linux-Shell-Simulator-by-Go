package services

import (
	"bytes"
	"projectshell/src/databases"
	command_model "projectshell/src/services/commands"
	session_user "projectshell/src/services/users"
	"testing"
)

func TestHistoryCommand(t *testing.T) {
	originalGetCurrentUser := session_user.GetCurrentUser
	session_user.GetCurrentUser = func() string { return "test_user" }
	defer func() { session_user.GetCurrentUser = originalGetCurrentUser }()

	databases.SetupTestDB(&command_model.Command{}, &session_user.User{})
	databases.SetDB(databases.GetTestDB())

	db := databases.GetTestDB()
	// mock user
	db.Create(&session_user.User{ID: 1, UserName: "test_user"})

	// mock history
	db.Create(&command_model.Command{Text: "ls", CreatedBy: 1})
	db.Create(&command_model.Command{Text: "cd /home", CreatedBy: 1})
	db.Create(&command_model.Command{Text: "ls", CreatedBy: 1})

	var outputWriter, errorWriter bytes.Buffer
	HistoryCommand([]string{}, &outputWriter, &errorWriter)

	expectedOutput := `|Command             |Count     |
--------------------------------
|ls                  |2         |
|cd /home            |1         |
`

	if outputWriter.String() != expectedOutput {
		t.Errorf("expected output:\n%s\nbut got:\n%s\n", expectedOutput, outputWriter.String())
	}

	if errorWriter.Len() > 0 {
		t.Errorf("expected no error output, but got: %s", errorWriter.String())
	}
}

package utils

import (
	"projectshell/src/databases"
	command_model "projectshell/src/services/commands"
	user_model "projectshell/src/services/users"
	"testing"
)

func TestGetUserId(t *testing.T) {

	databases.SetupTestDB(&user_model.User{})
	databases.SetDB(databases.GetTestDB())
	db := databases.GetTestDB()

	testUser := &user_model.User{UserName: "test_user"}
	db.Create(testUser)

	userID := GetUserId("test_user")
	if userID != testUser.ID {
		t.Errorf("expected user ID %d, got %d", testUser.ID, userID)
	}

	userID = GetUserId("non_existing_user")
	if userID != 0 {
		t.Errorf("expected user ID 0 for non-existing user, got %d", userID)
	}
}

func TestCleanHistory(t *testing.T) {
	databases.SetupTestDB(&user_model.User{}, &command_model.Command{})
	databases.SetDB(databases.GetTestDB())
	db := databases.GetTestDB()

	testUser := &user_model.User{UserName: "test_user"}
	db.Create(testUser)

	db.Create(&command_model.Command{Text: "ls", CreatedBy: testUser.ID})
	db.Create(&command_model.Command{Text: "cd /home", CreatedBy: testUser.ID})

	var count int64
	db.Model(&command_model.Command{}).Where("created_by = ?", testUser.ID).Count(&count)
	if count != 2 {
		t.Fatalf("expected 2 commands before clean, got %d", count)
	}

	err := CleanHistory(testUser.ID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	db.Model(&command_model.Command{}).Where("created_by = ?", testUser.ID).Count(&count)
	if count != 0 {
		t.Errorf("expected 0 commands after clean, got %d", count)
	}

	err = CleanHistory(9999)
	if err == nil || err.Error() != "user not found" {
		t.Errorf("expected 'user not found' error, got %v", err)
	}
}

package services

import (
	"fmt"
	"projectshell/src/databases"
	command_model "projectshell/src/services/commands"
	session_user "projectshell/src/services/users"
	"strings"
)

func LogHistory(args []string) {

	commandText := strings.Join(args, " ")

	currentUser := session_user.GetCurrentUser()
	if currentUser == "" {
		id := getUserId("anonymous")
		createHistory(id, commandText)
	} else {
		id := getUserId(currentUser)
		createHistory(id, commandText)
	}
}

func createHistory(id uint, commandText string) {

	db := databases.GetDB()
	newCommand := command_model.Command{
		Text:      commandText,
		CreatedBy: id,
	}

	if err := db.Create(&newCommand).Error; err != nil {
		fmt.Printf("Error creating command: %v\n", err)
		return
	}
}

func getUserId(username string) uint {

	db := databases.GetDB()
	userObj := &session_user.User{}
	if err := db.Where("user_name = ?", username).First(userObj).Error; err != nil {
		return 0
	}

	return userObj.ID
}

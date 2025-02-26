package services

import (
	"fmt"
	"projectshell/src/databases"
	command_model "projectshell/src/services/commands"
	session_user "projectshell/src/services/users"
	"projectshell/src/utils"
	"strings"
)

func LogHistory(args []string) {

	if args[0] == "history" {
		return
	}

	commandText := strings.Join(args, " ")

	currentUser := session_user.GetCurrentUser()
	if currentUser == "" {
		id := utils.GetUserId("anonymous")
		createHistory(id, commandText)
	} else {
		id := utils.GetUserId(currentUser)
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

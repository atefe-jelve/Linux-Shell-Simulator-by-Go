package services

import (
	"fmt"
	"projectshell/src/databases"
)

func LoginCommand(args []string) {

	if len(args) < 1 && args[0] == "clean"{
		fmt.Println("Not enough arguments provided.")
		return
	}

	var password_ string
	if len(args) > 1 {
		password_ = args[1]
	}

	username := args[0]
	password := password_

	ok, err := AuthenticateUser(username, password)
	if err != nil || !ok {
		fmt.Println("Invalid credentials")
		return
	}

	if !IsSessionValid(username) {
		AddUserSession(username)
	}
	SetCurrentUser(username)

}

func AuthenticateUser(username, password string) (bool, error) {
	db := databases.GetDB()
	user := &User{}

	if err := db.Where("user_name = ? AND password = ?", username, password).First(user).Error; err != nil {
		return false, err
	}

	return true, nil
}

package services

import (
	"fmt"
	"io"
	"projectshell/src/databases"
)

func LoginCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

	if len(args) < 1 || args[0] == "clean" {
		fmt.Fprintln(outputWriter, "Not enough arguments provided.")
		return
	}

	username := args[0]
	var password string
	if len(args) > 1 {
		password = args[1]
	}

	ok, err := AuthenticateUser(username, password)
	if err != nil || !ok {
		fmt.Fprintln(outputWriter, "Invalid credentials")
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

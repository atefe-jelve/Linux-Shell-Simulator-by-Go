package services

import (
	"fmt"
	"io"
	"projectshell/src/databases"

	"gorm.io/gorm"
)

func AddUserCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

	if len(args) == 0 || args[0] == "" {
		fmt.Fprintln(outputWriter, "Not enough arguments provided.")
		return
	}

	db := databases.GetDB()

	name := args[0]
	userObj := &User{}

	err := db.Where("user_name = ?", name).First(&userObj).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Fprintf(errorWriter, "Database error: %v\n", err)
		return
	}

	if err == nil {
		if args[0] != "anonymous" {
			fmt.Fprintln(errorWriter, "Duplicate user exists with this username.")
			return
		} else {
			// "anonymous" user exists — silently return (no error message needed per my logic).
			return
		}
	}

	var password_ string
	if len(args) > 1 {
		password_ = args[1]
	}

	newUser := User{
		UserName: name,
		Password: password_,
	}

	if err := db.Create(&newUser).Error; err != nil {
		fmt.Fprintf(errorWriter, "Error creating user: %v\n", err)
		return
	}

	fmt.Fprintln(outputWriter, "User created successfully.")
}

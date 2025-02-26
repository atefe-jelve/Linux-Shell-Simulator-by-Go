package services

import (
	"fmt"
	"projectshell/src/databases"

	"gorm.io/gorm"
)

func AddUserCommand(args []string) {
	if len(args) < 1 {
		// fmt.Println("Not enough arguments provided.")
		return
	}

	db := databases.GetDB()

	name := args[0]
	userObj := &User{}

	err := db.Where("user_name = ?", name).First(&userObj).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if err == nil {
		if args[0] != "anonymous" {
			fmt.Println("duplicate user exists with this username")
			return
		} else {
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
		fmt.Printf("Error creating user: %v\n", err)
		return
	}

	fmt.Println("user created successfully")
}

package utils

import (
	"errors"
	"projectshell/src/databases"
	command_ "projectshell/src/services/commands"
	user_ "projectshell/src/services/users"
)

func CleanHistory(id uint) error {

	db := databases.GetDB()
	if db == nil {
		return errors.New("unexpected database errors")
	}

	result := db.Where("created_by = ?", id).Delete(&command_.Command{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func GetUserId(username string) uint {

	db := databases.GetDB()
	userObj := &user_.User{}
	if err := db.Where("user_name = ?", username).First(userObj).Error; err != nil {
		return 0
	}

	return userObj.ID
}

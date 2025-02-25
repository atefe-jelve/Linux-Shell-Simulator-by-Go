package services

import (
	"errors"
	"projectshell/src/databases"
	command_model "projectshell/src/services/commands"
)

func CleanHistory(id uint) error {

	db := databases.GetDB()
	if db == nil {
		return errors.New("unexpected database errors")
	}

	result := db.Where("created_by = ?", id).Delete(&command_model.Command{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

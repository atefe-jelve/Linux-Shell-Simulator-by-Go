package services

import "time"

type Command struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Text      string
	CreatedBy uint      `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (Command) TableName() string {
	return "commands_history"
}

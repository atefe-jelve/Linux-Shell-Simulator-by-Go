package services

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserName  string    `gorm:"column:user_name"`
	Password  string    `gorm:"size:10 column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (User) TableName() string {
	return "users_shell"
}

package user

import "time"

type UserModel struct {
	ID uint64 `gorm:"primaryKey"`

	Username string `gorm:"size:50;uniqueIndex"`
	Email    string `gorm:"size:100;uniqueIndex"`
	Password string `gorm:"size:255"`

	Nickname string `gorm:"size:50"`
	Avatar   string `gorm:"size:255"`

	Role   string `gorm:"size:20"`
	Status int8   `gorm:"default:1"`

	LastLoginAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}

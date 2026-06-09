package user

import "time"

type User struct {
	ID uint64

	Username string
	Email    string
	Password string

	Nickname string
	Avatar   string

	Role string

	Status int8

	LastLoginAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

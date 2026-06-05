package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error

	GetByID(ctx context.Context, id uint64) (*User, error)

	GetByUsername(
		ctx context.Context,
		username string,
	) (*User, error)

	GetByEmail(
		ctx context.Context,
		email string,
	) (*User, error)

	Update(ctx context.Context, user *User) error

	Delete(ctx context.Context, id uint64) error
}

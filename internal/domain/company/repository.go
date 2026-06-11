package company

import "context"

type Repository interface {
	Create(
		ctx context.Context,
		company *Company,
	) error

	GetByID(
		ctx context.Context,
		id uint64,
	) (*Company, error)

	List(
		ctx context.Context,
	) ([]*Company, error)

	Update(
		ctx context.Context,
		company *Company,
	) error

	Delete(
		ctx context.Context,
		id uint64,
	) error
}

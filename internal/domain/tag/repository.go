package tag

import "context"

type Repository interface {
	Create(
		ctx context.Context,
		tag *Tag,
	) error

	GetByID(
		ctx context.Context,
		id uint64,
	) (*Tag, error)

	List(
		ctx context.Context,
	) ([]*Tag, error)

	Update(
		ctx context.Context,
		tag *Tag,
	) error

	Delete(
		ctx context.Context,
		id uint64,
	) error
}

package favorite

import "context"

type ListQuery struct {
	UserID uint64

	Offset int
	Limit  int

	Keyword string
}

type Repository interface {
	Create(
		ctx context.Context,
		favorite *Favorite,
	) error

	GetByIDAndUserID(
		ctx context.Context,
		id uint64,
		userID uint64,
	) (*Favorite, error)

	GetByUserIDAndGameID(
		ctx context.Context,
		userID uint64,
		gameID uint64,
	) (*Favorite, error)

	Update(
		ctx context.Context,
		favorite *Favorite,
	) error

	Delete(
		ctx context.Context,
		id uint64,
		userID uint64,
	) error

	List(
		ctx context.Context,
		query ListQuery,
	) ([]*Favorite, int64, error)
}

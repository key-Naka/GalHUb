package game

import "context"

type Repository interface {
	Create(
		ctx context.Context,
		game *Game,
	) error

	GetByID(
		ctx context.Context,
		id uint64,
	) (*Game, error)

	Update(
		ctx context.Context,
		game *Game,
	) error

	Delete(
		ctx context.Context,
		id uint64,
	) error

	List(
		ctx context.Context,
		offset int,
		limit int,
	) ([]*Game, error)
}

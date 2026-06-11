package gametag

import "context"

type Repository interface {
	ReplaceGameTags(
		ctx context.Context,
		gameID uint64,
		tagIDs []uint64,
	) error

	GetTagIDsByGameID(
		ctx context.Context,
		gameID uint64,
	) ([]uint64, error)
}

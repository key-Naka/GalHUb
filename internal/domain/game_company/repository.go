package game_company

import "context"

type Repository interface {
	ReplaceGameCompanies(
		ctx context.Context,
		gameID uint64,
		companyIDs []uint64,
	) error

	GetCompanyIDsByGameID(
		ctx context.Context,
		gameID uint64,
	) ([]uint64, error)
}

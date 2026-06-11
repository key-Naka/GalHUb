package game_company

import (
	"context"
	"galhub/internal/app/game_company/command"
	gamecompany "galhub/internal/domain/game_company"
)

type Service struct {
	repo gamecompany.Repository
}

func NewService(
	repo gamecompany.Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ReplaceGameCompanies(
	ctx context.Context,
	cmd command.ReplaceGameCompanies,
) error {
	return s.repo.ReplaceGameCompanies(
		ctx,
		cmd.GameID,
		cmd.CompanyIDs,
	)
}

func (s *Service) GetCompanyIDsByGameID(
	ctx context.Context,
	gameID uint64,
) ([]uint64, error) {
	return s.repo.GetCompanyIDsByGameID(
		ctx,
		gameID,
	)
}

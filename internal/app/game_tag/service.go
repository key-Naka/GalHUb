package game_tag

import (
	"context"
	"galhub/internal/app/game_tag/command"
	gametag "galhub/internal/domain/game_tage"
)

type Service struct {
	repo gametag.Repository
}

func NewService(
	repo gametag.Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ReplaceGameTags(
	ctx context.Context,
	cmd command.ReplaceGameTags,
) error {
	return s.repo.ReplaceGameTags(
		ctx,
		cmd.GameID,
		cmd.TagIDs,
	)
}

func (s *Service) GetTagIDsByGameID(
	ctx context.Context,
	gameID uint64,
) ([]uint64, error) {
	return s.repo.GetTagIDsByGameID(
		ctx,
		gameID,
	)
}

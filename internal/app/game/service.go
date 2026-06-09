package game

import (
	"context"
	"galhub/internal/app/game/command"
	"galhub/internal/domain/game"
	domain "galhub/internal/domain/game"
)

type Service struct {
	repo game.Repository
}

func NewService(
	repo game.Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) Create(
	ctx context.Context,
	cmd command.CreateGame,
) error {

	game := &domain.Game{
		Title:         cmd.Title,
		OriginalTitle: cmd.OriginalTitle,

		Cover: cmd.Cover,

		Description: cmd.Description,

		ReleaseDate: cmd.ReleaseDate,

		Status: 1,
	}

	return s.repo.Create(
		ctx,
		game,
	)
}

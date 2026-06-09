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
func (s *Service) GetByID(
	ctx context.Context,
	id uint64,
) (*domain.Game, error) {
	game, err := s.repo.GetByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, ErrGameNotFound
	}

	return game, nil
}
func (s *Service) List(
	ctx context.Context,
	page int,
	size int,
) ([]*domain.Game, error) {
	if page < 1 || size < 1 {
		return nil, ErrInvalidPagination
	}

	offset := (page - 1) * size

	return s.repo.List(
		ctx,
		offset,
		size,
	)
}

func (s *Service) Update(
	ctx context.Context,
	cmd command.UpdateGame,
) error {

	game, err := s.repo.GetByID(
		ctx,
		cmd.ID,
	)

	if err != nil {
		return err
	}

	if game == nil {
		return ErrGameNotFound
	}

	game.Title = cmd.Title
	game.OriginalTitle = cmd.OriginalTitle
	game.Cover = cmd.Cover
	game.Description = cmd.Description
	game.ReleaseDate = cmd.ReleaseDate
	game.Status = cmd.Status

	return s.repo.Update(
		ctx,
		game,
	)
}
func (s *Service) Delete(
	ctx context.Context,
	id uint64,
) error {
	game, err := s.repo.GetByID(
		ctx,
		id,
	)

	if err != nil {
		return err
	}

	if game == nil {
		return ErrGameNotFound
	}

	return s.repo.Delete(
		ctx,
		id,
	)
}

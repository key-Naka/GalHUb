package tag

import (
	"context"
	"galhub/internal/app/tag/command"
	domain "galhub/internal/domain/tag"
)

type Service struct {
	repo domain.Repository
}

func NewService(
	repo domain.Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(
	ctx context.Context,
	cmd command.CreateTag,
) error {

	tag := &domain.Tag{
		Name: cmd.Name,
	}

	return s.repo.Create(
		ctx,
		tag,
	)
}

func (s *Service) GetByID(
	ctx context.Context,
	id uint64,
) (*domain.Tag, error) {
	tag, err := s.repo.GetByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, err
	}

	if tag == nil {
		return nil, ErrTagNotFound
	}

	return tag, nil
}

func (s *Service) List(
	ctx context.Context,
) ([]*domain.Tag, error) {
	return s.repo.List(
		ctx,
	)
}

func (s *Service) Update(
	ctx context.Context,
	cmd command.UpdateTag,
) error {

	tag, err := s.repo.GetByID(
		ctx,
		cmd.ID,
	)

	if err != nil {
		return err
	}

	if tag == nil {
		return ErrTagNotFound
	}

	tag.Name = cmd.Name

	return s.repo.Update(
		ctx,
		tag,
	)
}

func (s *Service) Delete(
	ctx context.Context,
	id uint64,
) error {
	tag, err := s.repo.GetByID(
		ctx,
		id,
	)

	if err != nil {
		return err
	}

	if tag == nil {
		return ErrTagNotFound
	}

	return s.repo.Delete(
		ctx,
		id,
	)
}

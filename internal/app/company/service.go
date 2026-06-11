package company

import (
	"context"
	"galhub/internal/app/company/command"
	domain "galhub/internal/domain/company"
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
	cmd command.CreateCompany,
) error {

	company := &domain.Company{
		Name: cmd.Name,
	}

	return s.repo.Create(
		ctx,
		company,
	)
}

func (s *Service) GetByID(
	ctx context.Context,
	id uint64,
) (*domain.Company, error) {
	company, err := s.repo.GetByID(
		ctx,
		id,
	)

	if err != nil {
		return nil, err
	}

	if company == nil {
		return nil, ErrCompanyNotFound
	}

	return company, nil
}

func (s *Service) List(
	ctx context.Context,
) ([]*domain.Company, error) {
	return s.repo.List(
		ctx,
	)
}

func (s *Service) Update(
	ctx context.Context,
	cmd command.UpdateCompany,
) error {

	company, err := s.repo.GetByID(
		ctx,
		cmd.ID,
	)

	if err != nil {
		return err
	}

	if company == nil {
		return ErrCompanyNotFound
	}

	company.Name = cmd.Name

	return s.repo.Update(
		ctx,
		company,
	)
}

func (s *Service) Delete(
	ctx context.Context,
	id uint64,
) error {
	company, err := s.repo.GetByID(
		ctx,
		id,
	)

	if err != nil {
		return err
	}

	if company == nil {
		return ErrCompanyNotFound
	}

	return s.repo.Delete(
		ctx,
		id,
	)
}

package game

import (
	"context"
	"galhub/internal/app/game/command"
	"galhub/internal/app/game/query"
	domainCompany "galhub/internal/domain/company"
	"galhub/internal/domain/game"
	domain "galhub/internal/domain/game"
	domainGameCompany "galhub/internal/domain/game_company"
	domainGameTag "galhub/internal/domain/game_tage"
	domainTag "galhub/internal/domain/tag"
)

type Service struct {
	repo            game.Repository
	tagRepo         domainTag.Repository
	gameTagRepo     domainGameTag.Repository
	companyRepo     domainCompany.Repository
	gameCompanyRepo domainGameCompany.Repository
}

func NewService(
	repo game.Repository,
	tagRepo domainTag.Repository,
	gameTagRepo domainGameTag.Repository,
	companyRepo domainCompany.Repository,
	gameCompanyRepo domainGameCompany.Repository,
) *Service {
	return &Service{
		repo:            repo,
		tagRepo:         tagRepo,
		gameTagRepo:     gameTagRepo,
		companyRepo:     companyRepo,
		gameCompanyRepo: gameCompanyRepo,
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
func (s *Service) GetDetail(
	ctx context.Context,
	id uint64,
) (*query.DetailGame, error) {
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

	detail := &query.DetailGame{
		ID: game.ID,

		Title:         game.Title,
		OriginalTitle: game.OriginalTitle,

		Cover:       game.Cover,
		Description: game.Description,

		ReleaseDate: game.ReleaseDate,

		ViewCount:     game.ViewCount,
		FavoriteCount: game.FavoriteCount,

		Status: game.Status,

		Tags:      make([]*query.DetailTag, 0),
		Companies: make([]*query.DetailCompany, 0),

		CreatedAt: game.CreatedAt,
		UpdatedAt: game.UpdatedAt,
	}

	tagIDs, err := s.gameTagRepo.GetTagIDsByGameID(
		ctx,
		id,
	)

	if err != nil {
		return nil, err
	}

	if len(tagIDs) > 0 {
		detail.Tags = make([]*query.DetailTag, 0, len(tagIDs))

		for _, tagID := range tagIDs {
			tag, err := s.tagRepo.GetByID(
				ctx,
				tagID,
			)

			if err != nil {
				return nil, err
			}

			if tag == nil {
				continue
			}

			detail.Tags = append(
				detail.Tags,
				&query.DetailTag{
					ID: tag.ID,

					Name: tag.Name,

					CreatedAt: tag.CreatedAt,
					UpdatedAt: tag.UpdatedAt,
				},
			)
		}
	}

	companyIDs, err := s.gameCompanyRepo.GetCompanyIDsByGameID(
		ctx,
		id,
	)

	if err != nil {
		return nil, err
	}

	if len(companyIDs) > 0 {
		detail.Companies = make([]*query.DetailCompany, 0, len(companyIDs))

		for _, companyID := range companyIDs {
			company, err := s.companyRepo.GetByID(
				ctx,
				companyID,
			)

			if err != nil {
				return nil, err
			}

			if company == nil {
				continue
			}

			detail.Companies = append(
				detail.Companies,
				&query.DetailCompany{
					ID: company.ID,

					Name: company.Name,

					CreatedAt: company.CreatedAt,
					UpdatedAt: company.UpdatedAt,
				},
			)
		}
	}

	return detail, nil
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

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
	q query.ListGame,
) (*query.ListGameResult, error) {
	if q.Page < 1 || q.Size < 1 {
		return nil, ErrInvalidPagination
	}

	offset := (q.Page - 1) * q.Size

	games, total, err := s.repo.List(
		ctx,
		domain.ListQuery{
			Offset:    offset,
			Limit:     q.Size,
			Keyword:   q.Keyword,
			TagID:     q.TagID,
			CompanyID: q.CompanyID,
			Status:    q.Status,
		},
	)

	if err != nil {
		return nil, err
	}

	result := &query.ListGameResult{
		List:  make([]*query.ListGameItem, 0, len(games)),
		Total: total,
		Page:  q.Page,
		Size:  q.Size,
	}

	tagNameCache := make(map[uint64]string)
	companyNameCache := make(map[uint64]string)

	for _, game := range games {
		item := &query.ListGameItem{
			ID: game.ID,

			Title:         game.Title,
			OriginalTitle: game.OriginalTitle,

			Cover:       game.Cover,
			Description: game.Description,

			ReleaseDate: game.ReleaseDate,

			ViewCount:     game.ViewCount,
			FavoriteCount: game.FavoriteCount,

			Status: game.Status,

			Tags:      make([]string, 0),
			Companies: make([]string, 0),

			CreatedAt: game.CreatedAt,
			UpdatedAt: game.UpdatedAt,
		}

		tagIDs, err := s.gameTagRepo.GetTagIDsByGameID(
			ctx,
			game.ID,
		)

		if err != nil {
			return nil, err
		}

		for _, tagID := range tagIDs {
			tagName, ok := tagNameCache[tagID]

			if !ok {
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

				tagName = tag.Name
				tagNameCache[tagID] = tagName
			}

			item.Tags = append(
				item.Tags,
				tagName,
			)
		}

		companyIDs, err := s.gameCompanyRepo.GetCompanyIDsByGameID(
			ctx,
			game.ID,
		)

		if err != nil {
			return nil, err
		}

		for _, companyID := range companyIDs {
			companyName, ok := companyNameCache[companyID]

			if !ok {
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

				companyName = company.Name
				companyNameCache[companyID] = companyName
			}

			item.Companies = append(
				item.Companies,
				companyName,
			)
		}

		result.List = append(
			result.List,
			item,
		)
	}

	return result, nil
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

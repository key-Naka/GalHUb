package favorite

import (
	"context"
	"galhub/internal/app/favorite/command"
	"galhub/internal/app/favorite/query"
	domainCompany "galhub/internal/domain/company"
	domainFavorite "galhub/internal/domain/favorite"
	domainGame "galhub/internal/domain/game"
	domainGameCompany "galhub/internal/domain/game_company"
	domainGameTag "galhub/internal/domain/game_tage"
	domainTag "galhub/internal/domain/tag"
)

type Service struct {
	repo            domainFavorite.Repository
	gameRepo        domainGame.Repository
	tagRepo         domainTag.Repository
	gameTagRepo     domainGameTag.Repository
	companyRepo     domainCompany.Repository
	gameCompanyRepo domainGameCompany.Repository
}

func NewService(
	repo domainFavorite.Repository,
	gameRepo domainGame.Repository,
	tagRepo domainTag.Repository,
	gameTagRepo domainGameTag.Repository,
	companyRepo domainCompany.Repository,
	gameCompanyRepo domainGameCompany.Repository,
) *Service {
	return &Service{
		repo:            repo,
		gameRepo:        gameRepo,
		tagRepo:         tagRepo,
		gameTagRepo:     gameTagRepo,
		companyRepo:     companyRepo,
		gameCompanyRepo: gameCompanyRepo,
	}
}

func (s *Service) Create(
	ctx context.Context,
	cmd command.CreateFavorite,
) error {
	game, err := s.gameRepo.GetByID(
		ctx,
		cmd.GameID,
	)
	if err != nil {
		return err
	}
	if game == nil {
		return ErrGameNotFound
	}

	exists, err := s.repo.GetByUserIDAndGameID(
		ctx,
		cmd.UserID,
		cmd.GameID,
	)
	if err != nil {
		return err
	}
	if exists != nil {
		return ErrFavoriteExists
	}

	return s.repo.Create(
		ctx,
		&domainFavorite.Favorite{
			UserID: cmd.UserID,
			GameID: cmd.GameID,
		},
	)
}

func (s *Service) GetByID(
	ctx context.Context,
	id uint64,
	userID uint64,
) (*query.GetFavoriteResult, error) {
	favorite, err := s.repo.GetByIDAndUserID(
		ctx,
		id,
		userID,
	)
	if err != nil {
		return nil, err
	}
	if favorite == nil {
		return nil, ErrFavoriteNotFound
	}

	game, err := s.gameRepo.GetByID(
		ctx,
		favorite.GameID,
	)
	if err != nil {
		return nil, err
	}

	return s.buildGetFavoriteResult(
		ctx,
		favorite,
		game,
	)
}

func (s *Service) GetByGameID(
	ctx context.Context,
	userID uint64,
	gameID uint64,
) (*query.FavoriteStatusResult, error) {
	game, err := s.gameRepo.GetByID(
		ctx,
		gameID,
	)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, ErrGameNotFound
	}

	favorite, err := s.repo.GetByUserIDAndGameID(
		ctx,
		userID,
		gameID,
	)
	if err != nil {
		return nil, err
	}

	result := &query.FavoriteStatusResult{
		GameID:      gameID,
		IsFavorited: favorite != nil,
		FavoriteID:  nil,
	}

	if favorite != nil {
		result.FavoriteID = &favorite.ID
	}

	return result, nil
}

func (s *Service) List(
	ctx context.Context,
	q query.ListFavorite,
) (*query.ListFavoriteResult, error) {
	if q.Page < 1 || q.Size < 1 {
		return nil, ErrInvalidPagination
	}

	offset := (q.Page - 1) * q.Size

	favorites, total, err := s.repo.List(
		ctx,
		domainFavorite.ListQuery{
			UserID:  q.UserID,
			Offset:  offset,
			Limit:   q.Size,
			Keyword: q.Keyword,
		},
	)
	if err != nil {
		return nil, err
	}

	result := &query.ListFavoriteResult{
		List:  make([]*query.ListFavoriteItem, 0, len(favorites)),
		Total: total,
		Page:  q.Page,
		Size:  q.Size,
	}

	tagNameCache := make(map[uint64]string)
	companyNameCache := make(map[uint64]string)

	for _, favorite := range favorites {
		game, err := s.gameRepo.GetByID(
			ctx,
			favorite.GameID,
		)
		if err != nil {
			return nil, err
		}

		item, err := s.buildListFavoriteItem(
			ctx,
			favorite,
			game,
			tagNameCache,
			companyNameCache,
		)
		if err != nil {
			return nil, err
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
	cmd command.UpdateFavorite,
) error {
	favorite, err := s.repo.GetByIDAndUserID(
		ctx,
		cmd.ID,
		cmd.UserID,
	)
	if err != nil {
		return err
	}
	if favorite == nil {
		return ErrFavoriteNotFound
	}

	game, err := s.gameRepo.GetByID(
		ctx,
		cmd.GameID,
	)
	if err != nil {
		return err
	}
	if game == nil {
		return ErrGameNotFound
	}

	if favorite.GameID != cmd.GameID {
		exists, err := s.repo.GetByUserIDAndGameID(
			ctx,
			cmd.UserID,
			cmd.GameID,
		)
		if err != nil {
			return err
		}
		if exists != nil && exists.ID != favorite.ID {
			return ErrFavoriteExists
		}
	}

	favorite.GameID = cmd.GameID

	return s.repo.Update(
		ctx,
		favorite,
	)
}

func (s *Service) Delete(
	ctx context.Context,
	id uint64,
	userID uint64,
) error {
	favorite, err := s.repo.GetByIDAndUserID(
		ctx,
		id,
		userID,
	)
	if err != nil {
		return err
	}
	if favorite == nil {
		return ErrFavoriteNotFound
	}

	return s.repo.Delete(
		ctx,
		id,
		userID,
	)
}

func (s *Service) buildGetFavoriteResult(
	ctx context.Context,
	favorite *domainFavorite.Favorite,
	game *domainGame.Game,
) (*query.GetFavoriteResult, error) {
	result := &query.GetFavoriteResult{
		ID: favorite.ID,

		UserID: favorite.UserID,
		GameID: favorite.GameID,

		CreatedAt: favorite.CreatedAt,
	}

	if game != nil {
		result.GameTitle = game.Title
		result.GameOriginalTitle = game.OriginalTitle
		result.GameCover = game.Cover
		result.ReleaseDate = game.ReleaseDate
		result.GameStatus = game.Status
		result.FavoriteCount = game.FavoriteCount
	}

	tags, err := s.collectTagNames(
		ctx,
		favorite.GameID,
		make(map[uint64]string),
	)
	if err != nil {
		return nil, err
	}

	companies, err := s.collectCompanyNames(
		ctx,
		favorite.GameID,
		make(map[uint64]string),
	)
	if err != nil {
		return nil, err
	}

	result.Tags = tags
	result.Companies = companies

	return result, nil
}

func (s *Service) buildListFavoriteItem(
	ctx context.Context,
	favorite *domainFavorite.Favorite,
	game *domainGame.Game,
	tagNameCache map[uint64]string,
	companyNameCache map[uint64]string,
) (*query.ListFavoriteItem, error) {
	item := &query.ListFavoriteItem{
		ID: favorite.ID,

		UserID: favorite.UserID,
		GameID: favorite.GameID,

		CreatedAt: favorite.CreatedAt,
	}

	if game != nil {
		item.GameTitle = game.Title
		item.GameOriginalTitle = game.OriginalTitle
		item.GameCover = game.Cover
		item.ReleaseDate = game.ReleaseDate
		item.GameStatus = game.Status
		item.FavoriteCount = game.FavoriteCount
	}

	tags, err := s.collectTagNames(
		ctx,
		favorite.GameID,
		tagNameCache,
	)
	if err != nil {
		return nil, err
	}

	companies, err := s.collectCompanyNames(
		ctx,
		favorite.GameID,
		companyNameCache,
	)
	if err != nil {
		return nil, err
	}

	item.Tags = tags
	item.Companies = companies

	return item, nil
}

func (s *Service) collectTagNames(
	ctx context.Context,
	gameID uint64,
	tagNameCache map[uint64]string,
) ([]string, error) {
	tagIDs, err := s.gameTagRepo.GetTagIDsByGameID(
		ctx,
		gameID,
	)
	if err != nil {
		return nil, err
	}

	tags := make([]string, 0, len(tagIDs))

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

		tags = append(
			tags,
			tagName,
		)
	}

	return tags, nil
}

func (s *Service) collectCompanyNames(
	ctx context.Context,
	gameID uint64,
	companyNameCache map[uint64]string,
) ([]string, error) {
	companyIDs, err := s.gameCompanyRepo.GetCompanyIDsByGameID(
		ctx,
		gameID,
	)
	if err != nil {
		return nil, err
	}

	companies := make([]string, 0, len(companyIDs))

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

		companies = append(
			companies,
			companyName,
		)
	}

	return companies, nil
}

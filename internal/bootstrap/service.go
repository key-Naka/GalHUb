package bootstrap

import (
	appCompany "galhub/internal/app/company"
	appFavorite "galhub/internal/app/favorite"
	appGame "galhub/internal/app/game"
	appGameCompany "galhub/internal/app/game_company"
	appGameTag "galhub/internal/app/game_tag"
	appTag "galhub/internal/app/tag"
	appUser "galhub/internal/app/user"
	"galhub/internal/infrastructure/config"
)

type Services struct {
	User        *appUser.Service
	Game        *appGame.Service
	Favorite    *appFavorite.Service
	Tag         *appTag.Service
	GameTag     *appGameTag.Service
	Company     *appCompany.Service
	GameCompany *appGameCompany.Service
}

func NewServices(
	repos *Repositories,
	cfg *config.Config,
) *Services {
	jwtConfig := config.JWTConfig{}
	if cfg != nil {
		jwtConfig = cfg.JWT
	}

	return &Services{
		User: appUser.NewService(
			repos.User,

			jwtConfig,
		),
		Game: appGame.NewService(
			repos.Game,
			repos.Tag,
			repos.GameTag,
			repos.Company,
			repos.GameCompany,
		),
		Favorite: appFavorite.NewService(
			repos.Favorite,
			repos.Game,
			repos.Tag,
			repos.GameTag,
			repos.Company,
			repos.GameCompany,
		),
		Tag: appTag.NewService(
			repos.Tag,
		),
		GameTag: appGameTag.NewService(
			repos.GameTag,
		),
		Company: appCompany.NewService(
			repos.Company,
		),
		GameCompany: appGameCompany.NewService(
			repos.GameCompany,
		),
	}
}

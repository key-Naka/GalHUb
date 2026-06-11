package bootstrap

import (
	appGame "galhub/internal/app/game"
	appGameTag "galhub/internal/app/game_tag"
	appTag "galhub/internal/app/tag"
	appUser "galhub/internal/app/user"
	"galhub/internal/infrastructure/config"
)

type Services struct {
	User    *appUser.Service
	Game    *appGame.Service
	Tag     *appTag.Service
	GameTag *appGameTag.Service
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
		),
		Tag: appTag.NewService(
			repos.Tag,
		),
		GameTag: appGameTag.NewService(
			repos.GameTag,
		),
	}
}

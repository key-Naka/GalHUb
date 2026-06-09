package bootstrap

import (
	appGame "galhub/internal/app/game"
	appUser "galhub/internal/app/user"
	"galhub/internal/infrastructure/config"
)

type Services struct {
	User *appUser.Service
	Game *appGame.Service
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
	}
}

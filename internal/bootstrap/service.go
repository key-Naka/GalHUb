package bootstrap

import (
	appUser "galhub/internal/app/user"
	"galhub/internal/infrastructure/config"
)

type Services struct {
	User *appUser.Service
}

func NewServices(
	repos *Repositories,
	cfg *config.Config,
) *Services {

	return &Services{
		User: appUser.NewService(
			repos.User,
			cfg.JWT,
		),
	}
}

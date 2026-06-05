package bootstrap

import appUser "galhub/internal/app/user"

type Services struct {
	User *appUser.Service
}

func NewServices(
	repos *Repositories,
) *Services {

	return &Services{
		User: appUser.NewService(
			repos.User,
		),
	}
}

package bootstrap

import "galhub/internal/infrastructure/config"

type Bootstrap struct {
	Config *config.Config

	Repositories *Repositories

	Services *Services

	Handlers *Handlers
}

func New() *Bootstrap {

	repos := NewRepositories()

	services := NewServices(repos)

	handlers := NewHandlers(services)

	return &Bootstrap{
		Repositories: repos,
		Services:     services,
		Handlers:     handlers,
	}
}

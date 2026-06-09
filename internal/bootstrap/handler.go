package bootstrap

import (
	"galhub/internal/interfaces/http/handler"
)

type Handlers struct {
	User *handler.UserHandler
	Game *handler.GameHandler
}

func NewHandlers(
	services *Services,
) *Handlers {

	return &Handlers{
		User: handler.NewUserHandler(
			services.User,
		),
		Game: handler.NewGameHandler(
			services.Game,
		),
	}
}

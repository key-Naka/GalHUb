package bootstrap

import (
	"galhub/internal/interfaces/http/handler"
)

type Handlers struct {
	User *handler.UserHandler
}

func NewHandlers(
	services *Services,
) *Handlers {

	return &Handlers{
		User: handler.NewUserHandler(
			services.User,
		),
	}
}

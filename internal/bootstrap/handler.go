package bootstrap

import (
	"galhub/internal/interfaces/http/handler"
)

type Handlers struct {
	User    *handler.UserHandler
	Game    *handler.GameHandler
	Tag     *handler.TagHandler
	GameTag *handler.GameTagHandler
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
		Tag: handler.NewTagHandler(
			services.Tag,
		),
		GameTag: handler.NewGameTagHandler(
			services.GameTag,
		),
	}
}

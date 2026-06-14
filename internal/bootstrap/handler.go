package bootstrap

import (
	"galhub/internal/interfaces/http/handler"
)

type Handlers struct {
	User        *handler.UserHandler
	Game        *handler.GameHandler
	Favorite    *handler.FavoriteHandler
	Tag         *handler.TagHandler
	GameTag     *handler.GameTagHandler
	Company     *handler.CompanyHandler
	GameCompany *handler.GameCompanyHandler
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
		Favorite: handler.NewFavoriteHandler(
			services.Favorite,
		),
		Tag: handler.NewTagHandler(
			services.Tag,
		),
		GameTag: handler.NewGameTagHandler(
			services.GameTag,
		),
		Company: handler.NewCompanyHandler(
			services.Company,
		),
		GameCompany: handler.NewGameCompanyHandler(
			services.GameCompany,
		),
	}
}

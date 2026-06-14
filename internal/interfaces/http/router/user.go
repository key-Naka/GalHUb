package router

import (
	"galhub/internal/interfaces/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserPublic(
	group *gin.RouterGroup,
	handler *handler.UserHandler,
	favoriteHandler *handler.FavoriteHandler,
) {

	group.POST(
		"/register",
		handler.Register,
	)

	group.POST(
		"/login",
		handler.Login,
	)

	group.GET(
		"/:id/favorites",
		favoriteHandler.ListByUserID,
	)

}
func RegisterUserPrivate(
	group *gin.RouterGroup,
	handler *handler.UserHandler,
) {

	group.GET(
		"/profile",
		handler.Profile,
	)
	group.PUT(
		"/profile",
		handler.UpdateProfile,
	)
	group.PUT(
		"/password",
		handler.ChangePassword,
	)
}

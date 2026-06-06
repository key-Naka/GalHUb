package router

import (
	"galhub/internal/interfaces/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserPublic(
	group *gin.RouterGroup,
	handler *handler.UserHandler,
) {

	group.POST(
		"/register",
		handler.Register,
	)

	group.POST(
		"/login",
		handler.Login,
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
}

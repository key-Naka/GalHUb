package router

import (
	"galhub/internal/interfaces/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterGame(
	group *gin.RouterGroup,
	handler *handler.GameHandler,
) {

	group.POST(
		"",
		handler.Create,
	)
}

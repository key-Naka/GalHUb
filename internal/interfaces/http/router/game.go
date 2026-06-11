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
	group.GET("", handler.List)

	group.GET("/:id/detail", handler.GetDetail)
	group.GET("/:id", handler.GetByID)

	group.PUT("/:id", handler.Update)

	group.DELETE("/:id", handler.Delete)
}

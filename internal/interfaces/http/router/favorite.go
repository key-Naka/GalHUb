package router

import (
	"galhub/internal/interfaces/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterFavorite(
	group *gin.RouterGroup,
	handler *handler.FavoriteHandler,
) {
	group.POST("", handler.Create)
	group.GET("", handler.List)
	group.GET("/game/:game_id", handler.GetByGameID)
	group.GET("/:id", handler.GetByID)
	group.PUT("/:id", handler.Update)
	group.DELETE("/:id", handler.Delete)
}

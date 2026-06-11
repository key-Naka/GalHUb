package router

import (
	"galhub/internal/interfaces/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterTag(
	group *gin.RouterGroup,
	handler *handler.TagHandler,
) {
	group.POST("", handler.Create)
	group.GET("", handler.List)
	group.GET("/:id", handler.GetByID)
	group.PUT("/:id", handler.Update)
	group.DELETE("/:id", handler.Delete)
}

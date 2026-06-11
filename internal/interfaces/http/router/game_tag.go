package router

import (
	"galhub/internal/interfaces/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterGameTag(
	group *gin.RouterGroup,
	handler *handler.GameTagHandler,
) {
	group.GET("/:id/tags", handler.GetTags)
	group.PUT("/:id/tags", handler.ReplaceTags)
}

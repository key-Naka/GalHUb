package router

import (
	"galhub/internal/interfaces/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCompany(
	group *gin.RouterGroup,
	handler *handler.CompanyHandler,
) {
	group.POST("", handler.Create)
	group.GET("", handler.List)
	group.GET("/:id", handler.GetByID)
	group.PUT("/:id", handler.Update)
	group.DELETE("/:id", handler.Delete)
}

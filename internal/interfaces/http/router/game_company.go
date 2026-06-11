package router

import (
	"galhub/internal/interfaces/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterGameCompany(
	group *gin.RouterGroup,
	handler *handler.GameCompanyHandler,
) {
	group.GET("/:id/companies", handler.GetCompanies)
	group.PUT("/:id/companies", handler.ReplaceCompanies)
}

package router

import (
	app "galhub/internal/app/user"
	mysqlUser "galhub/internal/infrastructure/mysql/user"

	"galhub/internal/interfaces/http/handler"

	mysqlDB "galhub/internal/infrastructure/mysql"

	"github.com/gin-gonic/gin"
)

func RegisterUser(
	group *gin.RouterGroup,
) {

	repo := mysqlUser.NewRepository(
		mysqlDB.DB,
	)

	service := app.NewService(
		repo,
	)

	userHandler := handler.NewUserHandler(
		service,
	)

	group.POST("/register", userHandler.Register)
	group.POST("/login", userHandler.Login)
}

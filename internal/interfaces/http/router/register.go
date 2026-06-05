package router

import (
	"galhub/internal/interfaces/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPing(group *gin.RouterGroup) {

	pingHandler := handler.NewPingHandler()

	group.GET("/ping", pingHandler.Ping)
}

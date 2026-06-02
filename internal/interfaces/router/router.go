package router

import (
	"galhub/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	router := []RouteGroup{
		{
			Path:     "/api/v1",
			NeedAuth: false,
			RegisterFns: []func(*gin.RouterGroup){
				RegisterPing,
			},
		},
	}
	for _, route := range router {
		var group *gin.RouterGroup
		if route.NeedAuth {
			group = r.Group(route.Path, middleware.Auth())
		} else {
			group = r.Group(route.Path)
		}
		for _, fn := range route.RegisterFns {
			fn(group)
		}

	}
	return r
}

package router

import (
	"galhub/internal/bootstrap"
	"galhub/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(bs *bootstrap.Bootstrap) *gin.Engine {
	r := gin.Default()
	router := []RouteGroup{
		{
			Path:     "/api/v1",
			NeedAuth: false,
			RegisterFns: []func(*gin.RouterGroup){
				RegisterPing,
			},
		},
		{
			Path:     "/api/v1/users",
			NeedAuth: false,
			RegisterFns: []func(*gin.RouterGroup){
				func(g *gin.RouterGroup) {
					RegisterUser(g, bs.Handlers.User)
				},
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

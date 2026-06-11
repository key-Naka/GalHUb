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
					RegisterUserPublic(g, bs.Handlers.User)
				},
			},
		},
		{
			Path:     "/api/v1/users",
			NeedAuth: true,
			RegisterFns: []func(*gin.RouterGroup){
				func(g *gin.RouterGroup) {
					RegisterUserPrivate(g, bs.Handlers.User)
				},
			},
		},
		{
			Path:     "/api/v1/games",
			NeedAuth: true,
			RegisterFns: []func(*gin.RouterGroup){
				func(g *gin.RouterGroup) {
					RegisterGame(
						g,
						bs.Handlers.Game,
					)
					RegisterGameTag(
						g,
						bs.Handlers.GameTag,
					)
				},
			},
		},
		{
			Path:     "/api/v1/tags",
			NeedAuth: true,
			RegisterFns: []func(*gin.RouterGroup){
				func(g *gin.RouterGroup) {
					RegisterTag(
						g,
						bs.Handlers.Tag,
					)
				},
			},
		},
	}
	for _, route := range router {
		var group *gin.RouterGroup
		if route.NeedAuth {
			group = r.Group(route.Path, middleware.Auth(bs.Config.JWT.Secret))
		} else {
			group = r.Group(route.Path)
		}
		for _, fn := range route.RegisterFns {
			fn(group)
		}
	}
	return r
}

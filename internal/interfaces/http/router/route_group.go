package router

import "github.com/gin-gonic/gin"

type RouteGroup struct {
	Path        string
	NeedAuth    bool
	RegisterFns []func(*gin.RouterGroup)
}

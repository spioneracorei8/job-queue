package routes

import (
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

type route struct {
	g          *gin.Engine
	middleware middleware.MyMiddleware
}

func NewRoute(g *gin.Engine, middleware middleware.MyMiddleware) *route {
	return &route{
		g:          g,
		middleware: middleware,
	}
}

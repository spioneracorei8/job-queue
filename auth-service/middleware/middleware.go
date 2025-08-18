package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type MyMiddleware interface {
	ValidateRequiredHeader(key string) gin.HandlerFunc
	Logger() gin.HandlerFunc
}

type Middleware struct {
}

func InitMiddleware() MyMiddleware {
	return &Middleware{}
}

func (m *Middleware) Logger() gin.HandlerFunc {
	return func(g *gin.Context) {
		var method, path, ip = g.Request.Method, g.Request.URL.Path, g.ClientIP()
		g.Set("method", strings.ToUpper(method))
		g.Set("path", path)
		g.Set("ip", ip)
		g.Next()
	}
}

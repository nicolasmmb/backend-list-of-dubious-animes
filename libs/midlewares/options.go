package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IgnoreOPTIONS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" || ctx.Request.Method == "options" {
			ctx.AbortWithStatus(http.StatusOK)
		}
		ctx.Next()
	}
}

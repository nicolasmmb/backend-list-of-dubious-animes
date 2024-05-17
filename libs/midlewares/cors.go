package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type HeadersConfig struct {
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	AllowCredentials bool
}

func Headers(
	config HeadersConfig,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// accessControlAllowOrigin := strings.Join(config.AllowOrigins, ",")

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", config.AllowOrigins)
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", config.AllowMethods)
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", config.AllowHeaders)
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(config.AllowCredentials))

		ctx.Next()
	}

}

package auth

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	g := r.Group("/api/auth")
	g.POST(ROUTE_AUTH_USER, ValidateCredentials)
	g.POST(ROUTE_AUTH_TOKEN_IS_VALID, TokenIsValid)
}

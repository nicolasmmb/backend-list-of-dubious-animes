package auth

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.POST(ROUTE_AUTH_USER, ValidateCredentials)
}

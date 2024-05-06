package user

import (
	"backend/libs/env"
	middlewares "backend/libs/midlewares"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	g := r.Group("")
	g.Use(
		middlewares.JwtValidate(env.Data.JWT_SECRET),
	)
	g.GET(ROUTE_USER_BY_ID, GetUserById)
	g.GET(ROUTE_USER_WITH_FILTER, GetUserWithFilter)
	g.POST(ROUTE_CREATE_USER, CreateUser)
	g.PUT(ROUTE_UPDATE_USER, UpdateUser)
	g.DELETE(ROUTE_DELETE_USER_BY_ID, DeleteUserById)
}

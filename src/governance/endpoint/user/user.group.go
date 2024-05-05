package user

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.GET(ROUTE_USER_BY_ID, GetUserById)
	r.GET(ROUTE_USER_WITH_FILTER, GetUserWithFilter)
	r.POST(ROUTE_CREATE_USER, CreateUser)
	r.PUT(ROUTE_UPDATE_USER, UpdateUser)
	r.DELETE(ROUTE_DELETE_USER_BY_ID, DeleteUserById)
}

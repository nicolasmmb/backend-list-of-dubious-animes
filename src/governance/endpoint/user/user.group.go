package user

import (
	"backend/libs/env"
	middlewares "backend/libs/midlewares"

	"github.com/gin-gonic/gin"
)

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}

func Routes(r *gin.Engine) {
	r.GET(ROUTE_SEND_NOTIFICATIONS, SendNotifications)
	r.GET(ROUTE_USER_NOTIFICATION, HeadersMiddleware(), UserNotificationsSSE)
	g := r.Group("/api/user")
	g.Use(
		middlewares.JwtValidate(env.Data.JWT_SECRET),
	)
	g.GET(ROUTE_USER_BY_ID, GetUserById)
	g.GET(ROUTE_USER_WITH_FILTER, GetUserWithFilter)
	g.POST(ROUTE_CREATE_USER, CreateUser)
	g.PUT(ROUTE_UPDATE_USER, UpdateUser)
	g.DELETE(ROUTE_DELETE_USER_BY_ID, DeleteUserById)
}

package user

import (
	"io"

	"github.com/gin-gonic/gin"
)

const ROUTE_USER_NOTIFICATION = "/api/user/notifications"

var TOTAL_NOTIFICATIONS = 0

var CHANNEL_NOTIFICATIONS = make(chan string, 100)

func UserNotificationsSSE(c *gin.Context) {

	c.Stream(func(writter io.Writer) bool {
		select {
		case <-c.Writer.CloseNotify():
			return false
		case <-c.Request.Context().Done():
			return false
		case text := <-CHANNEL_NOTIFICATIONS:
			c.SSEvent("message", gin.H{
				"text": text,
			},
			)
			return true
		}
	},
	)
}

const ROUTE_SEND_NOTIFICATIONS = "/api/user/notifications/send"

func SendNotifications(c *gin.Context) {
	text := c.Query("text")

	CHANNEL_NOTIFICATIONS <- text

	c.JSON(200, gin.H{
		"message": "Notification sent",
		"text":    text,
	})
}

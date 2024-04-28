package main

import (
	"backend/libs/database/postgresql"
	userCmd "backend/src/governance/command/user"
	user "backend/src/governance/endpoint/user"
	userSrv "backend/src/governance/service/user"
	"log"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/niko-labs/libs-go/bus"
	helpers "github.com/niko-labs/libs-go/helper"
)

func init() {
	helpers.LoadEnv()
	postgresql.Connect()
	LoadBusHandlers()

}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/user", user.CreateUser)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(GetServerAddr())
}

func GetServerAddr() string {
	SERVER_PORT := os.Getenv("SERVER_PORT")
	SERVER_HOST := os.Getenv("SERVER_HOST")
	SERVER_ADDR := SERVER_HOST + ":" + SERVER_PORT
	log.Println("Server running on: ", SERVER_ADDR)
	return SERVER_ADDR
}

func LoadBusHandlers() {
	bus := bus.GetGlobal()
	_ = bus.RegisterCommandHandler(userCmd.CommandCreateUser{}, userSrv.CommandCreateUser)

}

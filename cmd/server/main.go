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
	r := gin.Default()

	r.POST(user.ROUTE_CREATE_USER, user.CreateUser)
	r.PUT(user.ROUTE_UPDATE_USER, user.UpdateUser)
	r.GET(user.ROUTE_USER_BY_ID, user.GetUserById)
	r.DELETE(user.ROUTE_DELETE_USER_BY_ID, user.DeleteUserById)

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
	_ = bus.RegisterCommandHandler(userCmd.CommandUpdateUser{}, userSrv.CommandUpdateUser)
	_ = bus.RegisterCommandHandler(userCmd.CommandGetUserById{}, userSrv.CommandGetUserById)
	_ = bus.RegisterCommandHandler(userCmd.CommandDeleteUserById{}, userSrv.CommandDeleteUserById)
}

package main

import (
	"backend/libs/database/postgresql"
	userCmd "backend/src/governance/command/user"
	user "backend/src/governance/endpoint/user"
	userSrv "backend/src/governance/service/user"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"context"
	"log"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/niko-labs/libs-go/bus"
	helpers "github.com/niko-labs/libs-go/helper"
	"github.com/niko-labs/libs-go/helper/opentel"
)

func init() {

	helpers.LoadEnv()
	postgresql.Connect()
	LoadBusHandlers()

}

func main() {
	baseCtx := context.Background()
	err, exp := opentel.InitTracer(
		opentel.NewTraceConfig(
			"backend-x",
			os.Getenv("OTLP_ENDPOINT"),
			"https://opentelemetry.io/schemas/1.24.0",
			os.Getenv("SERVICE_NAME"),
			os.Getenv("SERVICE_VERSION"),
			os.Getenv("SERVICE_NAMESPACE"),
			os.Getenv("DEPLOYMENT_ENVIRONMENT"),
		),
	)
	if err != nil {
		log.Fatalf("failed to create OTLP exporter: %v", err)
	}
	defer exp.Shutdown(baseCtx)

	r := gin.Default()
	r.Use(otelgin.Middleware("backend-x"))

	r.GET(user.ROUTE_USER_BY_ID, user.GetUserById)
	r.GET(user.ROUTE_USER_WITH_FILTER, user.GetUserWithFilter)
	r.POST(user.ROUTE_CREATE_USER, user.CreateUser)
	r.PUT(user.ROUTE_UPDATE_USER, user.UpdateUser)
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
	_ = bus.RegisterCommandHandler(userCmd.CommandGetUserWithFilter{}, userSrv.CommandGetUserWithFilter)
}

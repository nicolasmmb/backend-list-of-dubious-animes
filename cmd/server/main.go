package main

import (
	"backend/libs/database/postgresql"
	authCmd "backend/src/governance/command/auth"
	userCmd "backend/src/governance/command/user"
	auth "backend/src/governance/endpoint/auth"
	user "backend/src/governance/endpoint/user"
	authSrv "backend/src/governance/service/auth"
	userSrv "backend/src/governance/service/user"

	helpers "github.com/niko-labs/libs-go/helper"

	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/helper/middleware"
	"github.com/niko-labs/libs-go/helper/opentel"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
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
	r.Use(middleware.AddTraceIdHeader())

	auth.Routes(r)
	user.Routes(r)

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
	_ = bus.RegisterCommandHandler(authCmd.CommandAuthValidateCredentials{}, authSrv.CommandAuthValidateCredentials)
}

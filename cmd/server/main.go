package main

import (
	"backend/libs/database/postgresql"
	"backend/libs/env"
	middlewares "backend/libs/midlewares"
	authCmd "backend/src/governance/command/auth"
	userCmd "backend/src/governance/command/user"
	auth "backend/src/governance/endpoint/auth"
	user "backend/src/governance/endpoint/user"
	authSrv "backend/src/governance/service/auth"
	userSrv "backend/src/governance/service/user"

	helpers "github.com/niko-labs/libs-go/helper"

	"context"
	"log"

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
			env.Data.OTLP_ENDPOINT,
			env.Data.OTLP_SCHEMA_URL,
			env.Data.SERVICE_NAME,
			env.Data.SERVICE_VERSION,
			env.Data.SERVICE_NAMESPACE,
			env.Data.DEPLOYMENT_ENVIRONMENT,
		),
	)

	if err != nil {
		log.Fatalf("failed to create OTLP exporter: %v", err)
	}
	defer exp.Shutdown(baseCtx)

	r := gin.Default()
	r.Use(otelgin.Middleware("backend-x"))
	r.Use(middleware.AddTraceIdHeader())
	r.Use(middlewares.Headers(
		middlewares.HeadersConfig{
			AllowOrigins:     "*",
			AllowMethods:     "*",
			AllowHeaders:     "*",
			AllowCredentials: true,
		},
	))
	r.Use(middlewares.IgnoreOPTIONS())

	auth.Routes(r)
	user.Routes(r)

	log.Println("Server running on: ", env.Data.GetServerAddr())
	r.Run(env.Data.GetServerAddr())

}

func LoadBusHandlers() {
	bus := bus.GetGlobal()
	// user
	_ = bus.RegisterCommandHandler(userCmd.CommandCreateUser{}, userSrv.CommandCreateUser)
	_ = bus.RegisterCommandHandler(userCmd.CommandUpdateUser{}, userSrv.CommandUpdateUser)
	_ = bus.RegisterCommandHandler(userCmd.CommandGetUserById{}, userSrv.CommandGetUserById)
	_ = bus.RegisterCommandHandler(userCmd.CommandDeleteUserById{}, userSrv.CommandDeleteUserById)
	_ = bus.RegisterCommandHandler(userCmd.CommandGetUserWithFilter{}, userSrv.CommandGetUserWithFilter)
	// auth
	_ = bus.RegisterCommandHandler(authCmd.CommandTokenIsValid{}, authSrv.CommandTokenIsValid)
	_ = bus.RegisterCommandHandler(authCmd.CommandAuthValidateCredentials{}, authSrv.CommandAuthValidateCredentials)
}

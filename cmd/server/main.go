package main

import (
	"backend/libs/database/postgresql"
	"backend/libs/tracer"
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
	"go.opentelemetry.io/otel"
)

func init() {

	helpers.LoadEnv()
	postgresql.Connect()
	LoadBusHandlers()

}

func main() {
	baseCtx := context.Background()

	exp, err := tracer.NewOTLPExporter(baseCtx)
	// exp, err := tracer.NewConsoleExporter()

	if err != nil {
		log.Fatalf("failed to create OTLP exporter: %v", err)
	}

	tp := tracer.NewTraceProvider(exp)
	defer tp.Shutdown(baseCtx)

	otel.SetTracerProvider(tp)

	trcr := tp.Tracer("backend-x")
	tracer.SaveTracer(trcr)

	// _, span := trcr.Start(baseCtx, "main")
	// for i := 0; i < 7; i++ {
	// 	_, span := trcr.Start(baseCtx, "main"+fmt.Sprintf("%d", i))
	// 	span.AddEvent("Event in the loop")
	// 	span.End()
	// }
	// defer span.End()
	// span.End()

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

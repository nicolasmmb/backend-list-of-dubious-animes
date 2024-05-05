package auth

import (
	"backend/libs/database/postgresql"
	command "backend/src/governance/command/auth"

	"backend/src/governance/models/auth"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/helper/opentel"
	"github.com/niko-labs/libs-go/uow"
)

const ROUTE_AUTH_USER = "/auth"

func AutenticateUser(c *gin.Context) {

	var body auth.AuthUserModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := body.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := opentel.GetTracer()
	ctx, span := t.Start(c.Request.Context(), "route-autenticate-user")
	defer span.End()

	db := postgresql.GetConnection()
	uow := uow.NewUnitOfWorkWithOptions(db, uow.WithSchema("animes"), uow.WithTracer(t), uow.WithContext(ctx))
	bus := bus.GetGlobal()

	result, err := bus.SendCommand(c.Request.Context(), command.CommandAuthValidateCredentials{
		Email:    body.Email,
		Password: body.Password,
	}, uow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

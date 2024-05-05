package user

import (
	"backend/libs/database/postgresql"

	command "backend/src/governance/command/user"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/helper/opentel"
	"github.com/niko-labs/libs-go/helper/paginator"
	"github.com/niko-labs/libs-go/uow"
)

const ROUTE_USER_WITH_FILTER = "/user"

func GetUserWithFilter(c *gin.Context) {
	pageInfo, err := paginator.Create(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := opentel.GetTracer()
	ctx, span := t.Start(c.Request.Context(), "GetUserWithFilter")
	defer span.End()

	db := postgresql.GetConnection()
	uow := uow.NewUnitOfWorkWithOptions(db, uow.WithSchema("animes"), uow.WithTracer(t))
	bus := bus.GetGlobal()

	result, err := bus.SendCommand(ctx, command.CommandGetUserWithFilter{Pagination: *pageInfo}, uow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// need to cast the result to DTO where user is the entity
	c.JSON(http.StatusOK, result)
}

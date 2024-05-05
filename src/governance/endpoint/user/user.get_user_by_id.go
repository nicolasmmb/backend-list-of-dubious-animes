package user

import (
	"backend/libs/database/postgresql"
	"backend/libs/response"
	userCmd "backend/src/governance/command/user"
	userEntity "backend/src/governance/entitiy/user"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/helper/opentel"
	"github.com/niko-labs/libs-go/uow"
)

const ROUTE_USER_BY_ID = "/user/:id"

func GetUserById(c *gin.Context) {

	_id := c.Param("id")
	if _id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O id do usuário é obrigatório"})
		return
	}
	id, err := uuid.Parse(_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O id do usuário é inválido"})
		return
	}

	t := opentel.GetTracer()
	ctx, span := t.Start(c.Request.Context(), "GetUserById")
	defer span.End()

	db := postgresql.GetConnection()
	uow := uow.NewUnitOfWorkWithOptions(db, uow.WithSchema("animes"), uow.WithTracer(opentel.GetTracer()), uow.WithContext(ctx))
	bus := bus.GetGlobal()

	result, err := bus.SendCommand(ctx, userCmd.CommandGetUserById{ID: id}, uow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	r := result.(*userEntity.User)
	c.JSON(http.StatusOK, response.BaseResponse[userEntity.User]{
		Item: *r,
		Msg:  "Usuário encontrado com sucesso!",
	})
}

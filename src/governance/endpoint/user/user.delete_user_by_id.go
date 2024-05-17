package user

import (
	"backend/libs/database/postgresql"
	"backend/libs/response"
	userCmd "backend/src/governance/command/user"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/helper/opentel"
	"github.com/niko-labs/libs-go/uow"
)

const ROUTE_DELETE_USER_BY_ID = "/:id"

func DeleteUserById(c *gin.Context) {
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
	ctx, span := t.Start(c.Request.Context(), "DeleteUserById")
	defer span.End()

	db := postgresql.GetConnection()
	uow := uow.NewUnitOfWorkWithOptions(db, uow.WithSchema("animes"), uow.WithTracer(t))
	bus := bus.GetGlobal()

	result, err := bus.SendCommand(ctx, userCmd.CommandDeleteUserById{ID: id}, uow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	r := result.(*uuid.UUID)
	c.JSON(http.StatusOK, response.OnlyIdAndMsg{Msg: "O usuário foi deletado com sucesso!", ID: r.String()})
}

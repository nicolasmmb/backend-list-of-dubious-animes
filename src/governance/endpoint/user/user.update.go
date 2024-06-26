package user

import (
	"backend/libs/database/postgresql"
	"backend/libs/response"
	command "backend/src/governance/command/user"
	"backend/src/governance/models/user"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/helper/opentel"
	"github.com/niko-labs/libs-go/uow"
)

const ROUTE_UPDATE_USER = "/:id"

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O id do usuário é obrigatório"})
		return
	}
	var body user.UpdateUserModel
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
	ctx, span := t.Start(c.Request.Context(), "UpdateUser")
	defer span.End()

	db := postgresql.GetConnection()
	uow := uow.NewUnitOfWorkWithOptions(db, uow.WithSchema("animes"), uow.WithTracer(t))
	bus := bus.GetGlobal()

	result, err := bus.SendCommand(ctx, command.CommandUpdateUser{
		ID:       id,
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Avatar:   &body.Avatar,
	}, uow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r := result.(*uuid.UUID)
	c.JSON(http.StatusOK, response.OnlyIdAndMsg{Msg: "O usuário foi atualizado com sucesso!", ID: r.String()})
}

package user

import (
	"backend/libs/database/postgresql"
	"backend/libs/response"
	command "backend/src/governance/command/user"
	entity "backend/src/governance/entity/user"

	"backend/src/governance/models/user"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/helper/opentel"
	"github.com/niko-labs/libs-go/uow"
)

const ROUTE_CREATE_USER = "/"

func CreateUser(c *gin.Context) {

	var body user.CreateUserModel
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
	ctx, span := t.Start(c.Request.Context(), "CreateUser")
	defer span.End()

	db := postgresql.GetConnection()
	uow := uow.NewUnitOfWorkWithOptions(db, uow.WithSchema("animes"), uow.WithTracer(t), uow.WithContext(ctx))
	bus := bus.GetGlobal()

	result, err := bus.SendCommand(c.Request.Context(), command.CommandCreateUser{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Avatar:   &body.Avatar,
	}, uow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	r := result.(*entity.User)
	c.JSON(http.StatusOK, response.OnlyIdAndMsg{Msg: fmt.Sprintf("O usuário %s foi criado com sucesso!", r.Name), ID: r.ID.String()})
}

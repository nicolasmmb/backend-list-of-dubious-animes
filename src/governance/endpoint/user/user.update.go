package user

import (
	"backend/libs/database/postgresql"
	"backend/libs/response"
	userCmd "backend/src/governance/command/user"
	"backend/src/governance/dto/user"
	userEntity "backend/src/governance/entitiy/user"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/uow"
)

const ROUTE_UPDATE_USER = "/user/:id"

func UpdateUser(c *gin.Context) {
	var body user.UpdateUserDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := body.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := postgresql.GetConnection()
	uow := uow.NewUnitOfWorkWithOptions(uow.WithConnection(db), uow.WithSchema("animes"))
	bus := bus.GetGlobal()

	result, err := bus.SendCommand(c.Request.Context(), userCmd.CommandUpdateUser{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Avatar:   &body.Avatar,
	}, uow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r := result.(*userEntity.User)
	c.JSON(http.StatusOK, response.OnlyIdAndMsg{Msg: fmt.Sprintf("O usuário %s foi atualizado com sucesso!", r.Name), ID: r.ID.String()})
}

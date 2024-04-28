package user

import (
	"backend/libs/database/postgresql"
	userCmd "backend/src/governance/command/user"
	"backend/src/governance/dto/user"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/uow"
)

func CreateUser(c *gin.Context) {
	var userEntry user.UserEntry
	if err := c.ShouldBindJSON(&userEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := userEntry.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := postgresql.GetConnection()
	uow := uow.NewUnitOfWorkWithOptions(uow.WithConnection(db), uow.WithSchema("animes"))
	bus := bus.GetGlobal()

	result, err := bus.SendCommand(c.Request.Context(), userCmd.CommandCreateUser{
		Name:     userEntry.Name,
		Email:    userEntry.Email,
		Password: userEntry.Password,
		Avatar:   &userEntry.Avatar,
	}, uow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_ = result

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

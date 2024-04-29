package user

import (
	userCmd "backend/src/governance/command/user"

	"context"

	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/repository"
	"github.com/niko-labs/libs-go/uow"
)

func CommandGetUserById(ctx context.Context, uow *uow.UnitOfWork, cmd bus.CommandHandler) (data any, erro error) {
	cmdData := cmd.Data().(*userCmd.CommandGetUserById)

	repo := repository.NewRepositoryFromUoW(uow, &UserRepo)

	user, err := repo.Queries.GetUserByID(ctx, cmdData.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

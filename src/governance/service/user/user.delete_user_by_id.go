package user

import (
	command "backend/src/governance/command/user"
	"errors"

	"context"

	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/repository"
	"github.com/niko-labs/libs-go/uow"
)

func CommandDeleteUserById(ctx context.Context, uow *uow.UnitOfWork, cmd bus.CommandHandler) (data any, erro error) {
	cmdData := cmd.Data().(*command.CommandDeleteUserById)

	repo := repository.NewRepositoryFromUoW(uow, &UserRepo)

	user, err := repo.Queries.GetUserByID(ctx, cmdData.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("--> User not found")
	}

	if user.IsDeleted() {
		return nil, errors.New("--> User already deleted")
	}

	userId, err := repo.Queries.DeleteUserByID(ctx, cmdData.ID)
	if err != nil {
		return nil, err
	}

	return userId, nil
}

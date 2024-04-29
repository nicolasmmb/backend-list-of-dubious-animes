package user

import (
	userCmd "backend/src/governance/command/user"
	userEntity "backend/src/governance/entitiy/user"

	"context"

	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/repository"
	"github.com/niko-labs/libs-go/uow"
)

func CommandCreateUser(ctx context.Context, uow *uow.UnitOfWork, cmd bus.CommandHandler) (data any, erro error) {
	cmdData := cmd.Data().(*userCmd.CommandCreateUser)

	newUser, err := userEntity.NewInstance(cmdData.Name, cmdData.Email, cmdData.Password, cmdData.Avatar)
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepositoryFromUoW(uow, &UserRepo)

	newUser, err = repo.Queries.CreateNewUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

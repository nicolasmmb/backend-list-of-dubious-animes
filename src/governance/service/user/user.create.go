package user

import (
	command "backend/src/governance/command/user"
	entity "backend/src/governance/entity/user"

	"context"

	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/repository"
	"github.com/niko-labs/libs-go/uow"
)

func CommandCreateUser(ctx context.Context, uow *uow.UnitOfWork, cmd bus.CommandHandler) (data any, erro error) {
	cmdData := cmd.Data().(*command.CommandCreateUser)

	user_instance, err := entity.NewInstance(cmdData.Name, cmdData.Email, cmdData.Password, cmdData.Avatar)
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepositoryFromUoW(uow, &UserRepo)

	user_instance, err = repo.Queries.CreateNewUser(ctx, user_instance)
	if err != nil {
		return nil, err
	}

	return user_instance, nil
}

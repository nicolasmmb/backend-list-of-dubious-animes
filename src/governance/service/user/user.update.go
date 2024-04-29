package user

import (
	userCmd "backend/src/governance/command/user"
	userEntity "backend/src/governance/entitiy/user"

	"context"

	"github.com/google/uuid"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/repository"
	"github.com/niko-labs/libs-go/uow"
)

func CommandUpdateUser(ctx context.Context, uow *uow.UnitOfWork, cmd bus.CommandHandler) (data any, erro error) {
	cmdData := cmd.Data().(*userCmd.CommandUpdateUser)

	newUser, err := userEntity.NewInstance(
		cmdData.Name,
		cmdData.Email,
		cmdData.Password,
		cmdData.Avatar,
	)
	if err != nil {
		return nil, err
	}
	id := uuid.MustParse(cmdData.ID)

	repo := repository.NewRepositoryFromUoW(uow, &UserRepo)

	uId, err := repo.Queries.UpdateExistingUserById(ctx, id, newUser)
	if err != nil {
		return nil, err
	}

	return uId, nil
}

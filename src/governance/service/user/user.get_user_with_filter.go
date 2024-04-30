package user

import (
	userCmd "backend/src/governance/command/user"

	userModel "backend/src/governance/models/user"
	"context"

	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/helper/paginator"
	"github.com/niko-labs/libs-go/repository"
	"github.com/niko-labs/libs-go/uow"
)

func CommandGetUserWithFilter(ctx context.Context, uow *uow.UnitOfWork, cmd bus.CommandHandler) (data any, erro error) {
	cmdData := cmd.Data().(*userCmd.CommandGetUserWithFilter)

	repo := repository.NewRepositoryFromUoW(uow, &UserRepo)

	users, total, err := repo.Queries.GetUserWithFilter(ctx, cmdData.Pagination)
	if err != nil {
		return nil, err
	}

	var finalUsers []*userModel.BaseUserReturnModel
	for _, user := range users {
		finalUsers = append(finalUsers, userModel.ToBaseUserReturnModel(
			user.ID,
			user.Name,
			user.Email,
			user.Avatar,
			user.CreatedAt,
			user.UpdatedAt,
			user.DeletedAt,
		))
	}

	pagination := paginator.CreatePaginationResponse("", &cmdData.Pagination, total, finalUsers)

	return pagination, nil
}

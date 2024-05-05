package user

import (
	userCmd "backend/src/governance/command/user"
	"backend/src/governance/entitiy/user"

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
	trce := uow.GetTracer()
	_, span := trce.Start(ctx, "Parse 'USERS' to 'BaseUserReturnModel'")
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
	span.End()
	// delete users
	users = []*user.User{}
	_ = users

	_, span = trce.Start(ctx, "CreatePaginationResponse")
	pagination := paginator.CreatePaginationResponse("", &cmdData.Pagination, total, finalUsers)
	span.End()

	return pagination, nil
}

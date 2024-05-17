package auth

import (
	"backend/libs/env"
	command "backend/src/governance/command/auth"
	"backend/src/governance/models/auth"
	"time"

	er "backend/src/governance/error/auth"

	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/repository"
	"github.com/niko-labs/libs-go/uow"
	"go.opentelemetry.io/otel/attribute"
)

func CommandAuthValidateCredentials(ctx context.Context, uow *uow.UnitOfWork, cmd bus.CommandHandler) (data any, erro error) {
	trce := uow.GetTracer()
	_, span := trce.Start(ctx, "service.auth:ValidatingCredentials")
	defer span.End()

	cmdData := cmd.Data().(*command.CommandAuthValidateCredentials)
	repo := repository.NewRepositoryFromUoW(uow, &RepoAuth)

	user, err := repo.Queries.GetUserToValidateCredentials(ctx, cmdData.Email)
	if err != nil {
		span.AddEvent("Error getting user")
		return nil, err
	}
	if user == nil {
		span.AddEvent(er.ErrUserNotFound.Error())
		return nil, er.ErrUserNotFound
	}

	span.SetAttributes(attribute.KeyValue{Key: "user_id", Value: attribute.StringValue(user.ID.String())})
	if user.IsDeleted() {
		span.AddEvent(er.ErrUserIsDeleted.Error())
		return nil, er.ErrUserIsDeleted
	}
	passIsValid := user.HasValidPassword(cmdData.Password)
	if !passIsValid {
		span.AddEvent(er.ErrInvalidCredentials.Error())
		return nil, er.ErrInvalidCredentials
	}

	span.AddEvent("User validated")
	tk := auth.TokenModel{}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":   user.ID.String(),
			"email": user.Email,
			"exp":   time.Now().Add(time.Second * time.Duration(env.Data.JWT_EXPIRATION)).Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(env.Data.JWT_SECRET))
	if err != nil {
		span.AddEvent("Error generating token")
		return nil, err
	}
	tk.AccessToken = tokenString
	tk.RefreshToken = uuid.New().String()

	return tk, nil
}

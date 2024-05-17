package auth

import (
	"backend/libs/env"
	command "backend/src/governance/command/auth"
	"backend/src/governance/models/auth"

	"context"

	"github.com/golang-jwt/jwt/v5"

	"github.com/niko-labs/libs-go/bus"
	"github.com/niko-labs/libs-go/uow"
)

func CommandTokenIsValid(ctx context.Context, uow *uow.UnitOfWork, cmd bus.CommandHandler) (data any, erro error) {
	trce := uow.GetTracer()
	_, span := trce.Start(ctx, "service.auth:CommandTokenIsValid")
	defer span.End()

	cmdData := cmd.Data().(*command.CommandTokenIsValid)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(cmdData.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Data.JWT_SECRET), nil
	})
	tokenIsValid := err == nil

	tx := auth.TokenIsValidOutput{
		AccessToken: tokenIsValid,
	}

	return tx, nil
}

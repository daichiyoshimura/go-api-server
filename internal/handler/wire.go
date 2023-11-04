//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package handler

import (
	"awsomeapp/internal/module/account"

	"github.com/google/wire"
	"github.com/uptrace/bun"
)

func Wire(db *bun.DB) (*Handlers, error) {
	wire.Build(
		account.Wire,
		wire.Bind(new(IAccountUsecase), new(*account.AccountUsecase)),
		NewAccountHandler,
		wire.Struct(new(Handlers), "*"),
	)

	return &Handlers{}, nil
}

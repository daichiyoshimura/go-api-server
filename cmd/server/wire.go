//go:build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"awsomeapp/internal"
	"awsomeapp/internal/account"
	"awsomeapp/internal/repository"

	"github.com/google/wire"
	"github.com/uptrace/bun"
)

func Handlers(db *bun.DB) (*internal.Handlers, error) {
	wire.Build(
		repository.NewAccountRepository,
		wire.Bind(new(bun.IDB), new(*bun.DB)),
		wire.Bind(new(account.IAccountRepository), new(*repository.AccountRepository)),
		account.NewGetHandler,
		account.NewPostHandler,
		account.NewPutHandler,
		wire.Struct(new(internal.Handlers), "*"),
	)
	return &internal.Handlers{}, nil
}

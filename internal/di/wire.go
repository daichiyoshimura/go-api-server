//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"awsomeapp/internal/handler"
	"awsomeapp/internal/handler/account"
	"awsomeapp/internal/repository"

	"github.com/google/wire"
	"github.com/uptrace/bun"
)

func Wire(db *bun.DB) (*handler.Handlers, error) {

	wire.Build(
		repository.NewAccountRepository,
		wire.Bind(new(bun.IDB), new(*bun.DB)),
		account.NewAccountGetHandler,
		account.NewAccountPostHandler,
		account.NewAccountPutHandler,
		wire.Bind(new(account.IAccountRepository), new(*repository.AccountRepository)),
		wire.Struct(new(handler.Handlers), "*"),
	)
	return &handler.Handlers{}, nil
}

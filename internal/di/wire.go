//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"awsomeapp/internal"
	"awsomeapp/internal/account"
	"awsomeapp/internal/repository"

	"github.com/google/wire"
	"github.com/uptrace/bun"
)

func Wire(db *bun.DB) (*internal.Handlers, error) {

	wire.Build(
		repository.NewAccountRepository,
		wire.Bind(new(bun.IDB), new(*bun.DB)),
		account.NewGetHandler,
		account.NewPostHandler,
		account.NewPutHandler,
		wire.Bind(new(account.IAccountRepository), new(*repository.AccountRepository)),
		wire.Struct(new(internal.Handlers), "*"),
	)
	return &internal.Handlers{}, nil
}

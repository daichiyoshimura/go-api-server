//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"awsomeapp/internal/handler"
	"awsomeapp/internal/repository"

	"github.com/google/wire"
	"github.com/uptrace/bun"
)

func Wire(db *bun.DB) (*handler.Handlers, error) {

	wire.Build(
		repository.NewAccountRepository,
		wire.Bind(new(bun.IDB), new(*bun.DB)),
		handler.NewAccountHandler,
		wire.Bind(new(handler.IAccountRepository), new(*repository.AccountRepository)),
		wire.Struct(new(handler.Handlers), "*"),
	)
	return &handler.Handlers{}, nil
}

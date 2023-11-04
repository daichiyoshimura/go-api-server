//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package account

import (
	"awsomeapp/internal/module/account/internal/repository"

	"github.com/google/wire"
	"github.com/uptrace/bun"
)

func Wire(db *bun.DB) (*AccountUsecase, error) {
	wire.Build(
		wire.Bind(new(bun.IDB), new(*bun.DB)),
		repository.NewAccountRepository,
		wire.Bind(new(IAccountRepository), new(*repository.AccountRepository)),
		NewAccountUsecase,
	)

	return &AccountUsecase{}, nil
}
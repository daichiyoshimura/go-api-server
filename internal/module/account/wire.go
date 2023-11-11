//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package account

import (
	"awsomeapp/internal/module/account/internal/repository"
	"testing"

	"github.com/google/wire"
	"github.com/uptrace/bun"
	"go.uber.org/mock/gomock"
)

func Wire(db *bun.DB) (*AccountUsecase, error) {
	wire.Build(
		wire.Bind(new(bun.IDB), new(*bun.DB)),
		repository.NewAccountRepository,
		wire.Bind(new(iAccountRepository), new(*repository.AccountRepository)),
		NewAccountUsecase,
	)

	return &AccountUsecase{}, nil
}

func WireMock(t *testing.T) (*AccountUsecase, error) {
	wire.Build(
		wire.Bind(new(gomock.TestReporter), new(*testing.T)),
		provideMockController,
		NewMockiAccountRepository,
		wire.Bind(new(iAccountRepository), new(*MockiAccountRepository)),
		NewAccountUsecase,
	)

	return &AccountUsecase{}, nil
}

func provideMockController(t gomock.TestReporter) *gomock.Controller {
	return gomock.NewController(t)
}

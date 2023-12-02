//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package handler

import (
	"awsomeapp/internal/module/account"
	"testing"

	"github.com/google/wire"
	"github.com/uptrace/bun"
	gomock "go.uber.org/mock/gomock"
)

func Wire(db *bun.DB, jwtSigningKey []byte) (*Handlers, error) {
	wire.Build(
		account.Wire,
		wire.Bind(new(iAccountUsecase), new(*account.AccountUsecase)),
		NewAccountHandler,
		NewHealthHandler,
		NewAuthHandler,
		wire.Struct(new(Handlers), "*"),
	)

	return &Handlers{}, nil
}

func WireMock(t *testing.T, jwtSigningKey []byte) (*Handlers, error) {
	wire.Build(
		wire.Bind(new(gomock.TestReporter), new(*testing.T)),
		provideMockController,
		NewMockiAccountUsecase,
		wire.Bind(new(iAccountUsecase), new(*MockiAccountUsecase)),
		NewAccountHandler,
		NewHealthHandler,
		NewAuthHandler,
		wire.Struct(new(Handlers), "*"),
	)

	return &Handlers{}, nil
}

func provideMockController(t gomock.TestReporter) *gomock.Controller {
	return gomock.NewController(t)
}

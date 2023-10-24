package usecase

import (
	"awsomeapp/internal/account/model"
	"context"
)

type IAccountRepository interface {
	Insert(ctx context.Context, in *model.Account) (int64, error)
	FindByID(ctx context.Context, in *model.Account) (*model.Account, error)
	Update(ctx context.Context, in *model.Account) error
}

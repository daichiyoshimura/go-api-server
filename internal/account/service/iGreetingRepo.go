package service

import (
	"context"
	"awsomeapp/internal/account/model"
)

type IaccountRepository interface {
	Insert(ctx context.Context, in *model.Account) (int64, error)
	FindByID(ctx context.Context, in *model.Account) (*model.Account, error)
	Update(ctx context.Context, in *model.Account) error
}

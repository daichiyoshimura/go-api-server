package service

import (
	"context"
	"trygobun/internal/greeting/model"
)

type IGreetingRepository interface {
	Insert(ctx context.Context, in *model.Greeting) (int64, error)
	FindByID(ctx context.Context, in model.GreetingFindByIdInput) (*model.Greeting, error)
	FindByAccount(ctx context.Context, in model.GreetingFindByAccountInput) ([]model.Greeting, error)
}

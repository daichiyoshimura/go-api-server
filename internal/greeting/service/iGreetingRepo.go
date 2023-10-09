package service

import (
	"context"
	"trygobun/internal/greeting/model"
)

type IGreetingRepository interface {
	Insert(ctx context.Context, in *model.Greeting) (int64, error)
	FindByID(ctx context.Context, in *model.Greeting) (*model.Greeting, error)
	FindByAccount(ctx context.Context, in *model.Greeting) ([]model.Greeting, error)
	Update(ctx context.Context, in *model.Greeting) error
}

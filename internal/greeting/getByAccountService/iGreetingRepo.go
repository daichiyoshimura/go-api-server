package getByAccountService

import (
	"context"
	"trygobun/internal/greeting/model"
)

type IGreetingRepository interface {
	FindByAccount(ctx context.Context, in model.GreetingFindByAccountInput) ([]model.Greeting, error)
}

package getService

import (
	"context"
	"trygobun/internal/greeting/model"
)

type IGreetingRepository interface {
	FindByID(ctx context.Context, in model.GreetingFindByIdInput) (*model.Greeting, error)
}

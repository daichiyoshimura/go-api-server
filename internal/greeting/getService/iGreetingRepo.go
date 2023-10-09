package getService

import (
	"context"
	"trygobun/internal/greeting/model"
)

type IGreetingRepository interface {
	Find(ctx context.Context, in model.GreetingFindInput) (*model.Greeting, error)
}

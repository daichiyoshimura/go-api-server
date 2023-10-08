package getService

import (
	"context"
	"trygobun/internal/greeting/model"
)

type IGreetingRepository interface {
	Find(ctx context.Context, id uint) (*model.Greeting, error)
}

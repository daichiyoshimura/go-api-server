package registerService

import (
	"context"
	"trygobun/internal/greeting/model"
)

type IGreetingRepository interface {
	Insert(ctx context.Context, in *model.Greeting) (int64, error)
}

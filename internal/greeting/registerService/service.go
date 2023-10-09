package registerService

import (
	"context"
	"trygobun/internal/greeting/model"
)

type Service struct {
	repo IGreetingRepository
}

func NewService(repo IGreetingRepository) *Service {
	return &Service{
		repo: repo,
	}
}

type Input struct {
	AccountID int64
	Message string
}

type Output struct {
	ID int64
}

func (s *Service) Register(ctx context.Context, in *Input) (*Output, error) {
	id, err := s.repo.Insert(ctx, &model.Greeting{
		AccountID: in.AccountID,
		Message: in.Message,
	})
	if err != nil {
		return nil, err
	}
	return &Output{
		ID: id,
	}, nil
}

package getService

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
	ID int64
}

type Output struct {
	ID      int64
	Message string
}

func (s *Service) Get(ctx context.Context, in *Input) (*Output, error) {
	greeting, err := s.repo.FindByID(ctx, model.GreetingFindByIdInput{
		ID: in.ID,
	})
	if err != nil {
		return nil, err
	}
	return &Output{
		ID:      greeting.ID,
		Message: greeting.Message,
	}, err
}

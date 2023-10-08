package getService

import (
	"context"
)

type Service struct {
	repo IGreetingRepository
}

func NewService(repo IGreetingRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(ctx context.Context, in *Input) (*Output, error) {
	greeting, err := s.repo.Find(ctx, in.ID())
	if err != nil {
		return nil, err
	}

	return NewOutput(greeting.Message), nil
}

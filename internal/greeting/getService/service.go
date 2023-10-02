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
	s.repo.Find(ctx, in.ID)

	return &Output{
		Message: "OK",
	}, nil
}

package getByAccountService

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
}

type Output struct {
	Greetings []model.Greeting
}

func (s *Service) GetByAccount(ctx context.Context, in *Input) (*Output, error) {
	greetings, err := s.repo.FindByAccount(ctx, model.GreetingFindByAccountInput{
		AccountID: in.AccountID,
	})
	if err != nil {
		return nil, err
	}
	return &Output{
		Greetings: greetings,
	}, nil
}

package service

import (
	"context"
	"trygobun/internal/greeting/model"
)

type GetByAccountService struct {
	repo IGreetingRepository
}

func NewGetByAccountService(repo IGreetingRepository) *GetByAccountService {
	return &GetByAccountService{
		repo: repo,
	}
}

type GetByAccountInput struct {
	AccountID int64
}

type GetByAccountOutput struct {
	Greetings []model.Greeting
}

func (s *GetByAccountService) GetByAccount(ctx context.Context, in *GetByAccountInput) (*GetByAccountOutput, error) {
	greetings, err := s.repo.FindByAccount(ctx, &model.Greeting{
		AccountID: in.AccountID,
	})
	if err != nil {
		return nil, err
	}
	return &GetByAccountOutput{
		Greetings: greetings,
	}, nil
}

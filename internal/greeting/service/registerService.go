package service

import (
	"context"
	"trygobun/internal/greeting/model"
)

type RegisterService struct {
	repo IGreetingRepository
}

func NewRegisterService(repo IGreetingRepository) *RegisterService {
	return &RegisterService{
		repo: repo,
	}
}

type RegisterServiceInput struct {
	AccountID int64
	Message   string
}

type RegisterServiceOutput struct {
	ID int64
}

func (s *RegisterService) Register(ctx context.Context, in *RegisterServiceInput) (*RegisterServiceOutput, error) {
	id, err := s.repo.Insert(ctx, &model.Greeting{
		AccountID: in.AccountID,
		Message:   in.Message,
	})
	if err != nil {
		return nil, err
	}
	return &RegisterServiceOutput{
		ID: id,
	}, nil
}

package service

import (
	"context"
	"trygobun/internal/greeting/model"
)

type GetService struct {
	repo IGreetingRepository
}

func NewGetService(repo IGreetingRepository) *GetService {
	return &GetService{
		repo: repo,
	}
}

type GetServiceInput struct {
	ID int64
}

type GetServiceOutput struct {
	ID      int64
	Message string
}

func (s *GetService) Get(ctx context.Context, in *GetServiceInput) (*GetServiceOutput, error) {
	greeting, err := s.repo.FindByID(ctx, model.GreetingFindByIdInput{
		ID: in.ID,
	})
	if err != nil {
		return nil, err
	}
	return &GetServiceOutput{
		ID:      greeting.ID,
		Message: greeting.Message,
	}, err
}

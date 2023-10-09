package service

import (
	"context"
	"trygobun/internal/greeting/model"
)

type UpdateService struct {
	repo IGreetingRepository
}

func NewUpdateService(repo IGreetingRepository) *UpdateService {
	return &UpdateService{
		repo: repo,
	}
}

type UpdateServiceInput struct {
	ID        int64
	AccountID int64
	Message   string
}

type UpdateServiceOutput struct {
	Message string
}

func (s *UpdateService) Update(ctx context.Context, in *UpdateServiceInput) (*UpdateServiceOutput, error) {
	err := s.repo.Update(ctx, &model.Greeting{
		ID:        in.ID,
		AccountID: in.AccountID,
		Message:   in.Message,
	})
	if err != nil {
		return nil, err
	}
	return &UpdateServiceOutput{
		Message: "Affected",
	}, nil
}

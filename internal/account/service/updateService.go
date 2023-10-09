package service

import (
	"context"
	"trygobun/internal/account/model"
)

type UpdateService struct {
	repo IaccountRepository
}

func NewUpdateService(repo IaccountRepository) *UpdateService {
	return &UpdateService{
		repo: repo,
	}
}

type UpdateServiceInput struct {
	ID        int64
	Message   string
}

type UpdateServiceOutput struct {
	Message string
}

func (s *UpdateService) Update(ctx context.Context, in *UpdateServiceInput) (*UpdateServiceOutput, error) {
	err := s.repo.Update(ctx, &model.Account{
		ID:        in.ID,
		Message:   in.Message,
	})
	if err != nil {
		return nil, err
	}
	return &UpdateServiceOutput{
		Message: "Affected",
	}, nil
}

package service

import (
	"context"
	"trygobun/internal/account/model"
)

type GetService struct {
	repo IaccountRepository
}

func NewGetService(repo IaccountRepository) *GetService {
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
	account, err := s.repo.FindByID(ctx, &model.Account{
		ID: in.ID,
	})
	if err != nil {
		return nil, err
	}
	return &GetServiceOutput{
		ID:      account.ID,
		Message: account.Message,
	}, err
}

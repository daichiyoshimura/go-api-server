package service

import (
	"context"
	"trygobun/internal/account/model"
)

type RegisterService struct {
	repo IaccountRepository
}

func NewRegisterService(repo IaccountRepository) *RegisterService {
	return &RegisterService{
		repo: repo,
	}
}

type RegisterServiceInput struct {
	Message   string
}

type RegisterServiceOutput struct {
	ID int64
}

func (s *RegisterService) Register(ctx context.Context, in *RegisterServiceInput) (*RegisterServiceOutput, error) {
	id, err := s.repo.Insert(ctx, &model.Account{
		Message:   in.Message,
	})
	if err != nil {
		return nil, err
	}
	return &RegisterServiceOutput{
		ID: id,
	}, nil
}

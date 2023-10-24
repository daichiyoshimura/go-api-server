package service

import (
	"context"
	"awsomeapp/internal/account/model"
	"awsomeapp/internal/server"
)

type GetService struct {
	repo IaccountRepository
}

func NewGetService(repo IaccountRepository) *GetService {
	return &GetService{
		repo: repo,
	}
}

func (s *GetService) Get(ctx context.Context, in *server.Account) (*server.Account, error) {
	account, err := s.repo.FindByID(ctx, &model.Account{
		ID: in.Id,
	})
	if err != nil {
		return nil, err
	}
	return &server.Account{
		Id:   account.ID,
		Name: account.Name,
	}, err
}

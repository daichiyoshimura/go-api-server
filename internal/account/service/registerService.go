package service

import (
	"context"
	"awsomeapp/internal/account/model"
	"awsomeapp/internal/server"
)

type RegisterService struct {
	repo IaccountRepository
}

func NewRegisterService(repo IaccountRepository) *RegisterService {
	return &RegisterService{
		repo: repo,
	}
}

func (s *RegisterService) Register(ctx context.Context, in *server.NewAccount) (*server.Account, error) {
	id, err := s.repo.Insert(ctx, &model.Account{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	return &server.Account{
		Id: id,
	}, nil
}

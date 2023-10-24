package usecase

import (
	"awsomeapp/internal/account/model"
	"awsomeapp/internal/server"
	"context"
)

type CreateUsecase struct {
	repo IAccountRepository
}

func NewCreateUsecase(repo IAccountRepository) *CreateUsecase {
	return &CreateUsecase{
		repo: repo,
	}
}

func (s *CreateUsecase) Create(ctx context.Context, in *server.NewAccount) (*server.Account, error) {
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

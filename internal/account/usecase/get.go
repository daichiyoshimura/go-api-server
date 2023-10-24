package usecase

import (
	"awsomeapp/internal/account/model"
	"awsomeapp/internal/server"
	"context"
)

type GetUsecase struct {
	repo IAccountRepository
}

func NewGetUsecase(repo IAccountRepository) *GetUsecase {
	return &GetUsecase{
		repo: repo,
	}
}

func (s *GetUsecase) Get(ctx context.Context, in *server.Account) (*server.Account, error) {

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

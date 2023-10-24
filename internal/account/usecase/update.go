package usecase

import (
	"awsomeapp/internal/account/model"
	"awsomeapp/internal/server"
	"context"
)

type UpdateUsecase struct {
	repo IAccountRepository
}

func NewUpdateUsecase(repo IAccountRepository) *UpdateUsecase {
	return &UpdateUsecase{
		repo: repo,
	}
}

func (s *UpdateUsecase) Update(ctx context.Context, in *server.Account) (*server.Account, error) {
	err := s.repo.Update(ctx, &model.Account{
		ID:   in.Id,
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	return &server.Account{
		Name: "Affected",
	}, nil
}

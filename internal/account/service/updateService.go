package service

import (
	"context"
	"trygobun/internal/account/model"
	"trygobun/internal/server"
)

type UpdateService struct {
	repo IaccountRepository
}

func NewUpdateService(repo IaccountRepository) *UpdateService {
	return &UpdateService{
		repo: repo,
	}
}

func (s *UpdateService) Update(ctx context.Context, in *server.Account) (*server.Account, error) {
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

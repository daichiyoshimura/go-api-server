package getService

import "github.com/labstack/echo/v4"

type Service struct {
	repo IGreetingRepository
}

func NewService(repo IGreetingRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(ctx echo.Context, in *Input) (*Output, error) {
	s.repo.Find(ctx.Request().Context(), in.ID)

	return nil, nil
}

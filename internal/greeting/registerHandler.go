package greeting

import (
	"net/http"
	"trygobun/internal/greeting/repository"
	"trygobun/internal/greeting/service"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type RegisterRequest struct {
	AccountID int64  `json:"accountId"`
	Message   string `json:"message"`
}

type RegisterResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

func RegisterHandlerFunc(db bun.IDB) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var req RegisterRequest
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, &ErrorResponse{
				Message: err.Error(),
			})
		}

		repo := repository.NewGreetingRepository(db)
		out, err := service.NewRegisterService(repo).Register(ctx.Request().Context(), &service.RegisterServiceInput{
			AccountID: req.AccountID,
			Message:   req.Message,
		})
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, &RegisterResponse{
			ID: out.ID,
		})
	}
}

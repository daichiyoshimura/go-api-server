package greeting

import (
	"net/http"
	"trygobun/internal/greeting/registerService"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	AccountID int64  `json:"accountId"`
	Message   string `json:"message"`
}

type RegisterResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

func RegisterHandlerFunc(repo registerService.IGreetingRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var req RegisterRequest
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, &ErrorResponse{
				Message: err.Error(),
			})
		}

		out, err := registerService.NewService(repo).Register(ctx.Request().Context(), &registerService.Input{
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

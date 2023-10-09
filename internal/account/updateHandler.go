package account

import (
	"net/http"
	"trygobun/internal/account/repository"
	"trygobun/internal/account/service"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type UpdateRequest struct {
	AccountID int64  `json:"accountId"`
	Message   string `json:"message"`
}

type UpdateResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

func UpdateHandlerFunc(db bun.IDB) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var req UpdateRequest
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, &ErrorResponse{
				Message: err.Error(),
			})
		}

		repo := repository.NewAccountRepository(db)
		out, err := service.NewUpdateService(repo).Update(ctx.Request().Context(), &service.UpdateServiceInput{
			Message:   req.Message,
		})
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, &UpdateResponse{
			Message: out.Message,
		})
	}
}

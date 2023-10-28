package account

import (
	"awsomeapp/internal/account/usecase"
	"awsomeapp/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountGetHandler struct {
	repo IAccountRepository
}

func NewAccountGetHandler(repo IAccountRepository) *AccountGetHandler {
	return &AccountGetHandler{
		repo: repo,
	}
}

func (h *AccountGetHandler) GetAccount(ctx echo.Context, id int64) error {

	out, err := usecase.NewGetUsecase(h.repo).Get(ctx.Request().Context(), &server.Account{
		Id: id,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &server.Account{
		Id:   out.Id,
		Name: out.Name,
	})
}

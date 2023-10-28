package account

import (
	"awsomeapp/internal/account/usecase"
	"awsomeapp/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountPutHandler struct {
	repo IAccountRepository
}

func NewAccountPutHandler(repo IAccountRepository) *AccountPutHandler {
	return &AccountPutHandler{
		repo: repo,
	}
}

func (h *AccountPutHandler) PutAccount(ctx echo.Context, id int64) error {
	var req server.PutAccountJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &server.Error{
			Message: err.Error(),
		})
	}

	out, err := usecase.NewUpdateUsecase(h.repo).Update(ctx.Request().Context(), &server.Account{
		Name: req.Name,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &server.Account{
		Name: out.Name,
	})
}

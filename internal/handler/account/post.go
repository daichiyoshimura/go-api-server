package account

import (
	"awsomeapp/internal/account/usecase"
	"awsomeapp/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountPostHandler struct {
	repo IAccountRepository
}

func NewAccountPostHandler(repo IAccountRepository) *AccountPostHandler {
	return &AccountPostHandler{
		repo: repo,
	}
}

func (h *AccountPostHandler) PostAccount(ctx echo.Context) error {

	var req server.PostAccountJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &server.Error{
			Message: err.Error(),
		})
	}

	out, err := usecase.NewCreateUsecase(h.repo).Create(ctx.Request().Context(), &server.NewAccount{
		Name: req.Name,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &server.Account{
		Id: out.Id,
	})

}

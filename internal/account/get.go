package account

import (
	"awsomeapp/internal/account/usecase"
	"awsomeapp/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetHandler struct {
	repo IAccountRepository
}

func NewGetHandler(repo IAccountRepository) *GetHandler {
	return &GetHandler{
		repo: repo,
	}
}

func (h *GetHandler) GetAccount(ctx echo.Context, id int64) error {

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

package account

import (
	"awsomeapp/internal/account/repository"
	"awsomeapp/internal/account/usecase"
	"awsomeapp/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type PutHandler struct {
	db bun.IDB
}

func NewPutHandler(db bun.IDB) *PutHandler {
	return &PutHandler{
		db:db,
	}
}

func (h *PutHandler) PutAccount(ctx echo.Context, id int64) error {
	var req server.PutAccountJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &server.Error{
			Message: err.Error(),
		})
	}

	repo := repository.NewAccountRepository(h.db)
	out, err := usecase.NewUpdateUsecase(repo).Update(ctx.Request().Context(), &server.Account{
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

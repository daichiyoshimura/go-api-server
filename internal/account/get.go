package account

import (
	"awsomeapp/internal/account/repository"
	"awsomeapp/internal/account/usecase"
	"awsomeapp/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type GetHandler struct {
	db bun.IDB
}

func NewGetHandler(db bun.IDB) *GetHandler {
	return &GetHandler{
		db: db,
	}
}

func (h *GetHandler) GetAccount(ctx echo.Context, id int64) error {

	repo := repository.NewAccountRepository(h.db)
	out, err := usecase.NewGetUsecase(repo).Get(ctx.Request().Context(), &server.Account{
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
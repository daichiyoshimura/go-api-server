package account

import (
	"awsomeapp/internal/account/repository"
	"awsomeapp/internal/account/usecase"
	"awsomeapp/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type PostHandler struct {
	db bun.IDB
}

func NewPostHandler(db bun.IDB) *PostHandler {
	return &PostHandler{
		db: db,
	}
}

func (h *PostHandler) PostAccount(ctx echo.Context) error {

	var req server.PostAccountJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &server.Error{
			Message: err.Error(),
		})
	}

	repo := repository.NewAccountRepository(h.db)
	out, err := usecase.NewCreateUsecase(repo).Create(ctx.Request().Context(), &server.NewAccount{
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

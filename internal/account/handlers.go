package account

import (
	"awsomeapp/internal/account/repository"
	"awsomeapp/internal/account/usecase"
	"awsomeapp/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type Handlers struct {
	db bun.IDB
}

func NewHandlers(db bun.IDB) *Handlers {
	return &Handlers{
		db: db,
	}
}

func (h *Handlers) GetAccount(ctx echo.Context, id int64) error {

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

func (h *Handlers) PostAccount(ctx echo.Context) error {

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

func (h *Handlers) PutAccount(ctx echo.Context, id int64) error {
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

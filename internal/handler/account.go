package handler

import (
	"net/http"

	"awsomeapp/internal/server"
	"awsomeapp/internal/usecase"
	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	repo IAccountRepository
}

func NewAccountHandler(repo IAccountRepository) *AccountHandler {
	return &AccountHandler{
		repo: repo,
	}
}

func (h *AccountHandler) GetAccount(ctx echo.Context, id int64) error {
	out, err := usecase.NewAccountUsecase(h.repo).Get(&usecase.AccountGetInput{
		ID: id,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &server.Account{
		Id:   out.ID,
		Name: out.Name,
	})
}

func (h *AccountHandler) PostAccount(ctx echo.Context) error {
	var req server.PostAccountJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &server.Error{
			Message: err.Error(),
		})
	}

	out, err := usecase.NewAccountUsecase(h.repo).Create(&usecase.AccountCreateInput{
		Name: req.Name,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &server.Account{
		Id: out.ID,
	})
}

func (h *AccountHandler) PutAccount(ctx echo.Context, id int64) error {
	var req server.PutAccountJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &server.Error{
			Message: err.Error(),
		})
	}

	out, err := usecase.NewAccountUsecase(h.repo).Update(&usecase.AccountUpdateInput{
		ID:   req.Id,
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

func (h *AccountHandler) DeleteAccount(ctx echo.Context, id int64) error {
	err := usecase.NewAccountUsecase(h.repo).Delete(&usecase.AccountDeleteInput{
		ID: id,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	return ctx.NoContent(http.StatusOK)
}

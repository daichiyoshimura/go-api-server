package handler

import (
	"awsomeapp/internal/module/account"
	"awsomeapp/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	usecase iAccountUsecase
}

func NewAccountHandler(usecase iAccountUsecase) *AccountHandler {
	return &AccountHandler{
		usecase: usecase,
	}
}

func (h *AccountHandler) GetAccount(ctx echo.Context, id int64) error {
	out, err := h.usecase.Get(&account.AccountGetInput{
		ID: id,
	})

	if err != nil {
		writeLog(ctx, err)

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
		writeLog(ctx, err)

		return ctx.JSON(http.StatusBadRequest, &server.Error{
			Message: err.Error(),
		})
	}

	out, err := h.usecase.Create(&account.AccountCreateInput{
		Name: req.Name,
	})

	if err != nil {
		writeLog(ctx, err)

		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &server.Account{
		Id:   out.ID,
		Name: out.Name,
	})
}

func (h *AccountHandler) PutAccount(ctx echo.Context, id int64) error {
	var req server.PutAccountJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		writeLog(ctx, err)

		return ctx.JSON(http.StatusBadRequest, &server.Error{
			Message: err.Error(),
		})
	}

	out, err := h.usecase.Update(&account.AccountUpdateInput{
		ID:   req.Id,
		Name: req.Name,
	})

	if err != nil {
		writeLog(ctx, err)

		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &server.Account{
		Name: out.Name,
	})
}

func (h *AccountHandler) DeleteAccount(ctx echo.Context, id int64) error {
	err := h.usecase.Delete(&account.AccountDeleteInput{
		ID: id,
	})

	if err != nil {
		writeLog(ctx, err)

		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	return ctx.NoContent(http.StatusOK)
}

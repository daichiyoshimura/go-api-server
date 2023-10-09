package greeting

import (
	"net/http"
	"trygobun/internal/greeting/getByAccountService"

	"github.com/labstack/echo/v4"
)

type GetByAccountRequest struct {
	AccountID int64 `query:"accountId"`
}

type GetByAccountResponse struct {
	Greetings []greeting 
}

type greeting struct {
	Message string `json:"message"`
}

func GetByAccountHandlerFunc(repo getByAccountService.IGreetingRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var req GetByAccountRequest
		if err := ctx.Bind(&req); err != nil {
			return err
		}

		out, err := getByAccountService.NewService(repo).GetByAccount(ctx.Request().Context(), &getByAccountService.Input{
			AccountID: req.AccountID,
		})
		if err != nil {
			return err
		}

		greetings := []greeting{}
		for _,v := range out.Greetings {
			greetings = append(greetings, greeting{
				Message: v.Message,
			})
		}

		return ctx.JSON(http.StatusOK, &GetByAccountResponse{
			Greetings: greetings,
		})
	}
}

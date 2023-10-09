package greeting

import (
	"net/http"
	"trygobun/internal/greeting/repository"
	"trygobun/internal/greeting/service"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
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

func GetByAccountHandlerFunc(db bun.IDB) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var req GetByAccountRequest
		if err := ctx.Bind(&req); err != nil {
			return err
		}

		repo := repository.NewGreetingRepository(db)
		out, err := service.NewGetByAccountService(repo).GetByAccount(ctx.Request().Context(), &service.GetByAccountInput{
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

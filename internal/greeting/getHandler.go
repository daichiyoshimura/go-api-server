package greeting

import (
	"net/http"
	"trygobun/internal/greeting/getService"

	"github.com/labstack/echo/v4"
)

type GetRequest struct {
	ID int `param:"id"`
}
type GetResponse struct {
	Message string `json:"message"`
}

func GetHandlerFunc(repo getService.IGreetingRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var req GetRequest
		if err := ctx.Bind(&req); err != nil {
			return err
		}

		out, err := getService.NewService(repo).Get(ctx.Request().Context(), getService.NewInput(req.ID))
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, &GetResponse{
			Message: out.Message(),
		})
	}
}

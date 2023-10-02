package greeting

import (
	"net/http"
	"trygobun/internal/greeting/getService"

	"github.com/labstack/echo/v4"
)

type GetRequest struct {
	ID int
}
type GetResponse struct {
	Message string `json:"message"`
}

func GetHandlerFunc(repo getService.IGreetingRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		output, err := getService.NewService(repo).Get(ctx.Request().Context(), &getService.Input{
			ID: 1,
		})
		if err != nil {
			return err
		}

		r := &GetResponse{
			Message: output.Message,
		}

		return ctx.JSON(http.StatusOK, r)
	}
}

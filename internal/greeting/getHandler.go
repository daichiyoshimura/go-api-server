package greeting

import (
	"net/http"
	"trygobun/internal/greeting/getService"
	"trygobun/internal/greeting/repository"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func GetHandlerFunc(dbconn bun.IDB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		repo := repository.NewGreetingRepository(dbconn)
		srv := getService.NewService(repo)
		output, err := srv.Get(ctx, &getService.Input{
			ID: 1,
		})
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, &GetResponse{
			Message: output.Message,
		})
	}
}

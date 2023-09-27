package greeting

import (
	"net/http"
	"trygobun/internal/db"
	"trygobun/internal/greeting/getService"
	"trygobun/internal/greeting/repository"

	"github.com/labstack/echo/v4"
)

func GetHandler(ctx echo.Context) error {

	conn, err := db.Establish()
	if err != nil {
		return err
	}
	repo := repository.NewGreetingRepository(conn)
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

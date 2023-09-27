package greeting

import (
	"net/http"
	"trygobun/internal/db"
	"trygobun/internal/greeting/getOneService"
	"trygobun/internal/greeting/repository"

	"github.com/labstack/echo/v4"
)

func GetOneHandler(ctx echo.Context) error {

	conn, err := db.Establish()
	if err != nil {
		return err
	}
	repo := repository.NewGreetingRepository(conn)

	srv := getOneService.NewGreetingService(repo)
	output, err := srv.GetOne(&getOneService.Input{})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &GetOneResponse{
		Message: output.Message,
	})
}

package hello

import (
	"net/http"
	"trygobun/internal/dbdriver"

	"github.com/labstack/echo/v4"
)

func HelloHandler(ctx echo.Context) error {
	
	_, err := dbdriver.Establish()
	if err != nil {
		return err
	}

	


	return ctx.String(http.StatusOK, "Hello, World!")
}
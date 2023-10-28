package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func writeLog(ctx echo.Context, err error) {
	ctx.Logger().Errorj(log.JSON{
		"message":    err.Error(),
		"stackTrace": fmt.Sprintf("%+v", err),
	})
}

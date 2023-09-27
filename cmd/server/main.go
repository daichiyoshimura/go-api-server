package main

import (
	"trygobun/internal/greeting"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", greeting.GetHandler)
	e.Logger.Fatal(e.Start(":80"))
}
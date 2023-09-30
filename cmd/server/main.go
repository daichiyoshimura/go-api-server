package main

import (
	"trygobun/internal/db"
	"trygobun/internal/greeting"

	"github.com/labstack/echo/v4"
)

func main() {

	dbconn, err := db.NewConnection().Establish()
	if err != nil {
		panic(err)
	}
	
	e := echo.New()
	e.GET("/", greeting.GetHandlerFunc(dbconn))
	e.Logger.Fatal(e.Start(":80"))
}
package main

import (
	"trygobun/internal/db"
	"trygobun/internal/env"
	"trygobun/internal/greeting"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	srvEnv, dbEnv, err := env.NewReader().Read()
	if err != nil {
		e.Logger.Fatal(err)
	}

	dbconn, err := db.NewConnection().Establish(dbEnv)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", greeting.GetHandlerFunc(db.NewGreetingRepository(dbconn)))
	e.Logger.Fatal(e.Start(srvEnv.Host()))
}

package main

import (
	"trygobun/internal/db"
	"trygobun/internal/env"
	"trygobun/internal/greeting"

	"github.com/labstack/echo/v4"
)

func main() {

	srvEnv, dbEnv, err := env.NewReader().Read()
	if err != nil {
		panic(err)
	}

	dbconn, err := db.NewConnection().Establish(dbEnv)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/", greeting.GetHandlerFunc(dbconn))
	e.Logger.Fatal(e.Start(srvEnv.Host()))
}
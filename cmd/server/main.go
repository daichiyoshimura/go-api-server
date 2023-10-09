package main

import (
	"trygobun/internal/account"
	"trygobun/internal/db"
	"trygobun/internal/env"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	srvEnv, dbEnv, err := env.NewReader().Read()
	if err != nil {
		e.Logger.Fatal(err)
	}

	db, err := db.NewConnection().Establish(dbEnv)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/account/:id", account.GetHandlerFunc(db))
	e.PUT("/account/:id", account.UpdateHandlerFunc(db))
	e.POST("/account", account.RegisterHandlerFunc(db))

	if err := e.Start(srvEnv.Host()); err != nil {
		e.Logger.Fatal(err)
	}
}

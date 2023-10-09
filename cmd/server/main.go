package main

import (
	"trygobun/internal/account"
	"trygobun/internal/db"
	"trygobun/internal/env"
	"trygobun/internal/server"

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

	server.RegisterHandlers(e, account.NewHandlers(db))

	if err := e.Start(srvEnv.Host()); err != nil {
		e.Logger.Fatal(err)
	}
}

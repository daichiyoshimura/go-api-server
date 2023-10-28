package main

import (
	"awsomeapp/internal/db"
	"awsomeapp/internal/di"
	"awsomeapp/internal/env"
	"awsomeapp/internal/server"

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

	handlers, err := di.Wire(db)
	if err != nil {
		e.Logger.Fatal(err)
	}

	server.RegisterHandlersWithBaseURL(e, handlers, srvEnv.Host())

	if err := e.Start(srvEnv.Host()); err != nil {
		e.Logger.Fatal(err)
	}
}

package main

import (
	"awsomeapp/internal/db"
	"awsomeapp/internal/env"
	"awsomeapp/internal/handler"
	"awsomeapp/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	srvEnv, dbEnv, err := env.NewReader().Read()
	if err != nil {
		e.Logger.Fatal(err)
	}

	db, err := db.NewPool().Establish(dbEnv)
	if err != nil {
		e.Logger.Fatal(err)
	}

	handlers, err := handler.Wire(db)
	if err != nil {
		e.Logger.Fatal(err)
	}

	server.RegisterHandlersWithBaseURL(e, handlers, srvEnv.Host())

	if err := e.Start(srvEnv.Host()); err != nil {
		e.Logger.Fatal(err)
	}
}

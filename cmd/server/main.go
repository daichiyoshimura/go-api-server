package main

import (
	"awsomeapp/internal/db"
	"awsomeapp/internal/env"
	"awsomeapp/internal/handler"
	"awsomeapp/internal/log"
	"awsomeapp/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	logger := log.Logger()
	e.Use(log.RequestLogger(logger))

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

	server.RegisterHandlers(e, handlers)

	if err := e.Start(srvEnv.Host()); err != nil {
		e.Logger.Fatal(err)
	}
}

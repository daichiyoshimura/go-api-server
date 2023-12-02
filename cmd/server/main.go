package main

import (
	"awsomeapp/internal/auth"
	"awsomeapp/internal/db"
	"awsomeapp/internal/env"
	"awsomeapp/internal/handler"
	"awsomeapp/internal/log"
	"awsomeapp/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(log.RequestID())
	logger := log.Logger()
	e.Use(log.RequestLogger(logger))

	srvEnv, dbEnv, jwtEnv, err := env.NewReader().Read()
	if err != nil {
		e.Logger.Fatal(err)
	}

	signingKey := auth.SiginingKey(jwtEnv.Secret())
	e.Use(auth.JWT(signingKey))

	db, err := db.NewPool().Establish(dbEnv)
	if err != nil {
		e.Logger.Fatal(err)
	}

	handlers, err := handler.Wire(db, signingKey)
	if err != nil {
		e.Logger.Fatal(err)
	}

	server.RegisterHandlers(e, handlers)

	if err := e.Start(srvEnv.Host()); err != nil {
		e.Logger.Fatal(err)
	}
}

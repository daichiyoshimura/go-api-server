package main

import (
	"awsomeapp/internal/account"
	"awsomeapp/internal/db"
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

	server.RegisterHandlers(e, struct{
		*account.GetHandler
		*account.PostHandler
		*account.PutHandler
	}{
		account.NewGetHandler(db),
		account.NewPostHandler(db),
		account.NewPutHandler(db),
	})

	if err := e.Start(srvEnv.Host()); err != nil {
		e.Logger.Fatal(err)
	}
}

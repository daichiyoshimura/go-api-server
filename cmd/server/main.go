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

	db, err := db.NewConnection().Establish(dbEnv)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/greeting/:id", greeting.GetHandlerFunc(db))
	e.PUT("/greeting/:id", greeting.UpdateHandlerFunc(db))
	e.GET("/greeting", greeting.GetByAccountHandlerFunc(db))
	e.POST("/greeting", greeting.RegisterHandlerFunc(db))
	
	if err := e.Start(srvEnv.Host()); err != nil {
		e.Logger.Fatal(err)
	}
}

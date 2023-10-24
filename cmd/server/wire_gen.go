// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"awsomeapp/internal"
	"awsomeapp/internal/account"
	"awsomeapp/internal/repository"
	"github.com/uptrace/bun"
)

// Injectors from wire.go:

func Handlers(db *bun.DB) (*internal.Handlers, error) {
	accountRepository := repository.NewAccountRepository(db)
	getHandler := account.NewGetHandler(accountRepository)
	postHandler := account.NewPostHandler(accountRepository)
	putHandler := account.NewPutHandler(accountRepository)
	handlers := &internal.Handlers{
		GetHandler:  getHandler,
		PostHandler: postHandler,
		PutHandler:  putHandler,
	}
	return handlers, nil
}

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package api

// Account defines model for Account.
type Account struct {
	// Id Unique id of the account
	Id int64 `json:"id"`

	// Name Name of the account
	Name string `json:"name"`
}

// Error defines model for Error.
type Error struct {
	// Message Error message
	Message string `json:"message"`
}

// NewAccount defines model for NewAccount.
type NewAccount struct {
	// Name Name of the account
	Name string `json:"name"`
}

// PostAccountJSONRequestBody defines body for PostAccount for application/json ContentType.
type PostAccountJSONRequestBody = NewAccount

// PutAccountJSONRequestBody defines body for PutAccount for application/json ContentType.
type PutAccountJSONRequestBody = Account

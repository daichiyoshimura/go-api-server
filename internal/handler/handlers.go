package handler

import "awsomeapp/internal/handler/account"

type Handlers struct {
	*account.AccountGetHandler
	*account.AccountPostHandler
	*account.AccountPutHandler
}

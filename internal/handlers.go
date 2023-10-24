package internal

import "awsomeapp/internal/account"

type Handlers struct {
	*account.GetHandler
	*account.PostHandler
	*account.PutHandler
}

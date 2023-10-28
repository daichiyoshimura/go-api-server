package handler

import (
	"awsomeapp/internal/domain/account"
)

type IAccountRepository interface {
	Create(in *account.AccountDTO) (*account.AccountDTO, error)
	Get(id account.AccountID) (*account.AccountDTO, error)
	Update(in *account.AccountDTO) (*account.AccountDTO, error)
	Delete(id account.AccountID) error
}

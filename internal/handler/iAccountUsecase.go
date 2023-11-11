package handler

import "awsomeapp/internal/module/account"

type iAccountUsecase interface {
	Create(in *account.AccountCreateInput) (*account.Account, error)
	Get(in *account.AccountGetInput) (*account.Account, error)
	Update(in *account.AccountUpdateInput) (*account.Account, error)
	Delete(in *account.AccountDeleteInput) error
}

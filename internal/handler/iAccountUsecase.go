package handler

import "awsomeapp/internal/module/account"

type IAccountUsecase interface {
	Create(in *account.AccountCreateInput) (*account.AccountCreateOutput, error)
	Get(in *account.AccountGetInput) (*account.AccountGetOutput, error)
	Update(in *account.AccountUpdateInput) (*account.AccountUpdateOutput, error)
	Delete(in *account.AccountDeleteInput) error
}

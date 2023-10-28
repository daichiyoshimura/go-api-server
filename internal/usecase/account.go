package usecase

import (
	"awsomeapp/internal/domain/account"
)

type AccountUsecase struct {
	repo IAccountRepository
}

func NewAccountUsecase(repo IAccountRepository) *AccountUsecase {
	return &AccountUsecase{
		repo: repo,
	}
}

type Account struct {
	ID   int64
	Name string
}

func createAccountFromDTO(dto *account.AccountDTO) *Account {
	return &Account{
		ID:   int64(dto.ID),
		Name: string(dto.Name),
	}
}

type AccountCreateInput struct {
	Name string
}

type AccountCreateOutput Account

func (u *AccountUsecase) Create(in *AccountCreateInput) (*AccountCreateOutput, error) {

	ac, err := account.NewAccountService(u.repo).Create(&account.AccountDTO{
		Name: account.AccountName(in.Name),
	})
	if err != nil {
		return nil, err
	}
	return (*AccountCreateOutput)(createAccountFromDTO(ac.DTO())), nil
}

type AccountGetInput struct {
	ID int64
}

type AccountGetOutput Account

func (u *AccountUsecase) Get(in *AccountGetInput) (*AccountGetOutput, error) {

	ac, err := account.NewAccountService(u.repo).Get(account.AccountID(in.ID))
	if err != nil {
		return nil, err
	}

	return (*AccountGetOutput)(createAccountFromDTO(ac.DTO())), nil
}

type AccountUpdateInput Account

type AccountUpdateOutput Account

func (u *AccountUsecase) Update(in *AccountUpdateInput) (*AccountUpdateOutput, error) {
	ac, err := account.NewAccountService(u.repo).Update(&account.AccountDTO{
		ID:   account.AccountID(in.ID),
		Name: account.AccountName(in.Name),
	})
	if err != nil {
		return nil, err
	}

	return (*AccountUpdateOutput)(createAccountFromDTO(ac.DTO())), nil
}

type AccountDeleteInput struct {
	ID int64
}

func (u *AccountUsecase) Delete(in *AccountDeleteInput) error {
	return account.NewAccountService(u.repo).Delete(account.AccountID(in.ID))
}

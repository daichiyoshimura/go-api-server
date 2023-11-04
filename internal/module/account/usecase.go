package account

import (
	"awsomeapp/internal/module/account/internal/domain"

	"github.com/cockroachdb/errors"
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

func createAccountFromDTO(dto *domain.AccountDTO) *Account {
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
	ac, err := domain.NewAccountService(u.repo).Create(&domain.AccountCreateDTO{
		Name: domain.AccountName(in.Name),
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return (*AccountCreateOutput)(createAccountFromDTO(ac.DTO())), nil
}

type AccountGetInput struct {
	ID int64
}

type AccountGetOutput Account

func (u *AccountUsecase) Get(in *AccountGetInput) (*AccountGetOutput, error) {
	ac, err := domain.NewAccountService(u.repo).Get(domain.AccountID(in.ID))

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return (*AccountGetOutput)(createAccountFromDTO(ac.DTO())), nil
}

type AccountUpdateInput Account

type AccountUpdateOutput Account

func (u *AccountUsecase) Update(in *AccountUpdateInput) (*AccountUpdateOutput, error) {
	ac, err := domain.NewAccountService(u.repo).Update(&domain.AccountDTO{
		ID:   domain.AccountID(in.ID),
		Name: domain.AccountName(in.Name),
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return (*AccountUpdateOutput)(createAccountFromDTO(ac.DTO())), nil
}

type AccountDeleteInput struct {
	ID int64
}

func (u *AccountUsecase) Delete(in *AccountDeleteInput) error {
	if err := domain.NewAccountService(u.repo).Delete(domain.AccountID(in.ID)); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

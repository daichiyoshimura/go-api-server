package account

import (
	"awsomeapp/internal/module/account/internal/domain"

	"github.com/cockroachdb/errors"
)

type AccountUsecase struct {
	repo iAccountRepository
}

func NewAccountUsecase(repo iAccountRepository) *AccountUsecase {
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
		ID:   dto.ID,
		Name: dto.Name,
	}
}

type AccountCreateInput struct {
	Name string
}

func (i *AccountCreateInput) UnspecifiedDTO() *domain.AccountUnspecifiedDTO {
	return &domain.AccountUnspecifiedDTO{
		Name: i.Name,
	}
}

func (u *AccountUsecase) Create(in *AccountCreateInput) (*Account, error) {
	ac, err := domain.NewAccountService(u.repo).Create(in.UnspecifiedDTO())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return createAccountFromDTO(ac.DTO()), nil
}

type AccountGetInput struct {
	ID int64
}

func (u *AccountUsecase) Get(in *AccountGetInput) (*Account, error) {
	ac, err := domain.NewAccountService(u.repo).Get(in.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return createAccountFromDTO(ac.DTO()), nil
}

type AccountUpdateInput Account

func (i *AccountUpdateInput) DTO() *domain.AccountDTO {
	return &domain.AccountDTO{
		ID:   i.ID,
		Name: i.Name,
	}
}

func (u *AccountUsecase) Update(in *AccountUpdateInput) (*Account, error) {

	ac, err := domain.NewAccountService(u.repo).Update(in.DTO())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return createAccountFromDTO(ac.DTO()), nil
}

type AccountDeleteInput struct {
	ID int64
}

func (u *AccountUsecase) Delete(in *AccountDeleteInput) error {
	if err := domain.NewAccountService(u.repo).Delete(in.ID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

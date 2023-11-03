package account

import (
	"awsomeapp/internal/module/account/internal/domain"
	"awsomeapp/internal/module/account/internal/repository"

	"github.com/cockroachdb/errors"
	"github.com/uptrace/bun"
)

type AccountUsecase struct {
	db bun.IDB
}

func NewAccountUsecase(db bun.IDB) *AccountUsecase {
	return &AccountUsecase{
		db: db,
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
	accountRepo := repository.NewAccountRepository(u.db)
	ac, err := domain.NewAccountService(accountRepo).Create(&domain.AccountCreateDTO{
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
	accountRepo := repository.NewAccountRepository(u.db)
	ac, err := domain.NewAccountService(accountRepo).Get(domain.AccountID(in.ID))

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return (*AccountGetOutput)(createAccountFromDTO(ac.DTO())), nil
}

type AccountUpdateInput Account

type AccountUpdateOutput Account

func (u *AccountUsecase) Update(in *AccountUpdateInput) (*AccountUpdateOutput, error) {
	accountRepo := repository.NewAccountRepository(u.db)
	ac, err := domain.NewAccountService(accountRepo).Update(&domain.AccountDTO{
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
	accountRepo := repository.NewAccountRepository(u.db)
	if err := domain.NewAccountService(accountRepo).Delete(domain.AccountID(in.ID)); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

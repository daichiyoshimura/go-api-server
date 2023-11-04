package account

import "awsomeapp/internal/module/account/internal/domain"

type IAccountRepository interface {
	Create(in *domain.AccountCreateDTO) (*domain.AccountDTO, error)
	Get(id domain.AccountID) (*domain.AccountDTO, error)
	Update(in *domain.AccountDTO) (*domain.AccountDTO, error)
	Delete(id domain.AccountID) error
}

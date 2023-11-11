package account

import "awsomeapp/internal/module/account/internal/domain"

type iAccountRepository interface {
	Create(in *domain.AccountUnspecifiedDTO) (*domain.AccountDTO, error)
	Get(id int64) (*domain.AccountDTO, error)
	Update(in *domain.AccountDTO) (*domain.AccountDTO, error)
	Delete(id int64) error
}

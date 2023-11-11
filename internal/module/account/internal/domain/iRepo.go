package domain

type iAccountRepository interface {
	Create(udto *AccountUnspecifiedDTO) (*AccountDTO, error)
	Get(id int64) (*AccountDTO, error)
	Update(dto *AccountDTO) (*AccountDTO, error)
	Delete(id int64) error
}

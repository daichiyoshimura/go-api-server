package domain

type iAccountRepository interface {
	Create(udto *AccountUnspecifiedDTO) (*AccountDTO, error)
	Get(id string) (*AccountDTO, error)
	Update(dto *AccountDTO) (*AccountDTO, error)
	Delete(id string) error
}

package domain

type IAccountRepository interface {
	Create(in *AccountCreateDTO) (*AccountDTO, error)
	Get(id AccountID) (*AccountDTO, error)
	Update(in *AccountDTO) (*AccountDTO, error)
	Delete(id AccountID) error
}

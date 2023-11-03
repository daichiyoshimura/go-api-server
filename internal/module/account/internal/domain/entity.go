package domain

type AccountID int64

type AccountName string

type AccountEntity struct {
	id   AccountID
	name AccountName
}

type AccountDTO struct {
	ID   AccountID
	Name AccountName
}

type AccountCreateDTO struct {
	Name AccountName
}

func NewAccountEntity(dto *AccountDTO) *AccountEntity {
	return &AccountEntity{
		id:   dto.ID,
		name: dto.Name,
	}
}

func (e *AccountEntity) DTO() *AccountDTO {
	return &AccountDTO{
		ID:   e.id,
		Name: e.name,
	}
}

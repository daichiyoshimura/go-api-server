package domain

import "github.com/cockroachdb/errors"

type AccountID int64

type AccountName string

type AccountEntity struct {
	id   AccountID
	name AccountName
}

type AccountDTO struct {
	ID   *AccountID
	Name AccountName
}

func NewAccountEntity(dto *AccountDTO) (*AccountEntity, error) {
	if dto.ID == nil {
		return nil, errors.Newf("id must be set")
	}

	return &AccountEntity{
		id:   *dto.ID,
		name: dto.Name,
	}, nil
}

func (e *AccountEntity) DTO() *AccountDTO {
	return &AccountDTO{
		ID:   &e.id,
		Name: e.name,
	}
}

package domain

import (
	"awsomeapp/internal/module/account/internal/domain/value"

	"github.com/cockroachdb/errors"
)

type AccountEntity struct {
	id   string
	name *value.AccountName
}

type AccountDTO struct {
	ID   string
	Name string
}

type AccountUnspecifiedDTO struct {
	Name string
}

func NewAccountEntity(dto *AccountDTO) (*AccountEntity, error) {
	name, err := value.NewAccountName(dto.Name)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &AccountEntity{
		id:   dto.ID,
		name: name,
	}, nil
}

func (e *AccountEntity) DTO() *AccountDTO {
	return &AccountDTO{
		ID:   e.id,
		Name: e.name.Value(),
	}
}

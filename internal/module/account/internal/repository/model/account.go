package model

import (
	"time"

	"awsomeapp/internal/module/account/internal/domain"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts,alias:ac"`
	ID            []byte    `bun:"id,pk,type:binary(16)"`
	Name          string    `bun:"name,notnull"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `bun:"delete_at,soft_delete,nullzero"`
}

func CreateAccountFromID(id string) (*Account, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	binID, err := parsedID.MarshalBinary()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Account{
		ID: binID,
	}, nil
}

func CreateAccountFromDTO(dto *domain.AccountDTO) (*Account, error) {
	parsedID, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	binID, err := parsedID.MarshalBinary()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Account{
		ID:   binID,
		Name: dto.Name,
	}, nil
}

func CreateAccountFromUnspecifiedDTO(udto *domain.AccountUnspecifiedDTO) (*Account, error) {
	parsedID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	binID, err := parsedID.MarshalBinary()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Account{
		ID:   binID,
		Name: udto.Name,
	}, nil
}

func (a *Account) DTO() (*domain.AccountDTO, error) {
	id, err := uuid.ParseBytes(a.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &domain.AccountDTO{
		ID:   id.String(),
		Name: a.Name,
	}, nil
}

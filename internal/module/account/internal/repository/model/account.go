package model

import (
	"time"

	"awsomeapp/internal/module/account/internal/domain"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts"`
	ID            []byte    `bun:"id,pk"`
	Name          string    `bun:"name,notnull"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `bun:"delete_at,soft_delete,nullzero"`
}

func CreateAccountFromID(id string) (*Account, error) {
	binId, err := uuidToBin(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Account{
		ID: binId,
	}, nil
}

func CreateAccountFromDTO(dto *domain.AccountDTO) (*Account, error) {
	binId, err := uuidToBin(dto.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Account{
		ID:   binId,
		Name: dto.Name,
	}, nil
}

func CreateAccountFromUnspecifiedDTO(udto *domain.AccountUnspecifiedDTO) (*Account, error) {
	binId, err := generateUUID()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Account{
		ID:   binId,
		Name: udto.Name,
	}, nil
}

func uuidToBin(id string) ([]byte, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	bin, err := parsed.MarshalBinary()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return bin, nil
}

func uuidToStr(id []byte) (string, error) {
	parsed, err := uuid.ParseBytes(id)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return parsed.String(), nil
}

func generateUUID() ([]byte, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	binId, err := id.MarshalBinary()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return binId, nil
}

func (a *Account) DTO() (*domain.AccountDTO, error) {
	id, err := uuidToStr(a.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &domain.AccountDTO{
		ID:   id,
		Name: a.Name,
	}, nil
}

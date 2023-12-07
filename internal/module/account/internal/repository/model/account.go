package model

import (
	"time"

	"awsomeapp/internal/module/account/internal/domain"

	"github.com/cockroachdb/errors"

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

func (a *Account) DTO() (*domain.AccountDTO, error) {
	id, err := newUUIDHelper().toString(a.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &domain.AccountDTO{
		ID:   id,
		Name: a.Name,
	}, nil
}

type AccountFactory struct {
	uuidHelper iUUIDHelper
}

func NewAccountFactory(uuidHelper iUUIDHelper) *AccountFactory {
	return &AccountFactory{
		uuidHelper: uuidHelper,
	}
}

func (f *AccountFactory) CreateFromID(id string) (*Account, error) {
	binId, err := f.uuidHelper.toBinary(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Account{
		ID: binId,
	}, nil
}

func (f *AccountFactory) CreateFromDTO(dto *domain.AccountDTO) (*Account, error) {
	binId, err := newUUIDHelper().toBinary(dto.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Account{
		ID:   binId,
		Name: dto.Name,
	}, nil
}

func (f *AccountFactory) CreateAccountFromUnspecifiedDTO(udto *domain.AccountUnspecifiedDTO) (*Account, error) {
	binId, err := newUUIDHelper().generate()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Account{
		ID:   binId,
		Name: udto.Name,
	}, nil
}

func CreateAccountFromID(id string) (*Account, error) {
	binId, err := newUUIDHelper().toBinary(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Account{
		ID: binId,
	}, nil
}

func CreateAccountFromDTO(dto *domain.AccountDTO) (*Account, error) {
	binId, err := newUUIDHelper().toBinary(dto.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Account{
		ID:   binId,
		Name: dto.Name,
	}, nil
}

func CreateAccountFromUnspecifiedDTO(udto *domain.AccountUnspecifiedDTO) (*Account, error) {
	binId, err := newUUIDHelper().generate()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Account{
		ID:   binId,
		Name: udto.Name,
	}, nil
}

package model

import (
	"awsomeapp/internal/domain/account"
	"time"

	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts"`
	ID            int64     `bun:"id,pk"`
	Name          string    `bun:"name,notnull"`
	CreatedAt     time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,default:current_timestamp"`
	DeletedAt     time.Time `bun:",soft_delete,nullzero"`
}

func CreateAccountFromDTO(dto *account.AccountDTO) *Account {
	return &Account{
		ID:   int64(dto.ID),
		Name: string(dto.Name),
	}
}

func (a *Account) DTO() *account.AccountDTO {
	return &account.AccountDTO{
		ID:   account.AccountID(a.ID),
		Name: account.AccountName(a.Name),
	}
}

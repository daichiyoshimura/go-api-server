package model

import (
	"time"

	"awsomeapp/internal/module/account/internal/domain"

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

func CreateAccountFromDTO(dto *domain.AccountDTO) *Account {
	return &Account{
		ID:   int64(*dto.ID),
		Name: string(dto.Name),
	}
}

func (a *Account) DTO() *domain.AccountDTO {
	return &domain.AccountDTO{
		ID:   (*domain.AccountID)(&a.ID),
		Name: domain.AccountName(a.Name),
	}
}

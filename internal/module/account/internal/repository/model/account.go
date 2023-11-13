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
	CreatedAt     time.Time `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
	DeletedAt     time.Time `bun:"delete_at,soft_delete,nullzero"`
}

func CreateAccountFromDTO(dto *domain.AccountDTO) *Account {
	return &Account{
		ID:   dto.ID,
		Name: dto.Name,
	}
}

func CreateAccountFromUnspecifiedDTO(udto *domain.AccountUnspecifiedDTO) *Account {
	return &Account{
		Name: udto.Name,
	}
}

func (a *Account) DTO() *domain.AccountDTO {
	return &domain.AccountDTO{
		ID:   a.ID,
		Name: a.Name,
	}
}

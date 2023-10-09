package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts"`
	ID            int64     `bun:"id,pk"`
	Name          string    `bun:"name,notnull"`
	CreatedAt     time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,default:current_timestamp"`
}

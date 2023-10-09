package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Greeting struct {
	bun.BaseModel `bun:"table:greetings"`
	ID            int64     `bun:"id,pk"`
	AccountID     int64     `bun:"account_id,notnull"`
	Message       string    `bun:"message,notnull"`
	CreatedAt     time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,default:current_timestamp"`
}

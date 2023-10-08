package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Greeting struct {
	bun.BaseModel `bun:"table:greetings"`
	ID            int       `bun:"id,pk"`
	Message       string    `bun:"message,notnull"`
	CreatedAt     time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,default:current_timestamp"`
}

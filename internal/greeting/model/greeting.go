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

type GreetingFindByIdInput struct {
	ID int64
}

type GreetingFindByAccountInput struct {
	AccountID int64
}

type GreetingInsertInput struct {
	AccountID int64
	Message string
}

type GreetingUpdateInput struct {
	ID      int64
	Message string
}

package db

import (
	"context"
	"trygobun/internal/greeting/model"

	"github.com/uptrace/bun"
)

type GreetingRepository struct {
	conn bun.IDB
}

func NewGreetingRepository(conn bun.IDB) *GreetingRepository {
	return &GreetingRepository{
		conn: conn,
	}
}

func (r *GreetingRepository) Find(ctx context.Context, in model.GreetingFindInput) (*model.Greeting, error) {
	query := `
	SELECT 
		id,
		message
	FROM 
		greetings
	WHERE 
		id = ?
	`
	greetings := make([]model.Greeting, 1)
	if err := r.conn.NewRaw(query, greetings, in.ID).Scan(ctx, greetings); err != nil {
		return nil, err
	}
	return &greetings[0], nil
}

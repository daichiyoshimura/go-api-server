package repository

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

func (r *GreetingRepository) FindByID(ctx context.Context, in *model.Greeting) (*model.Greeting, error) {
	greetings := make([]model.Greeting, 1)
	if err := r.conn.NewSelect().Model(greetings).Where("id = ?", in.ID).Scan(ctx); err != nil {
		return nil, err
	}
	return &greetings[0], nil
}

func (r *GreetingRepository) FindByAccount(ctx context.Context, in *model.Greeting) ([]model.Greeting, error) {
	greetings := make([]model.Greeting, 2)
	if err := r.conn.NewSelect().Model(greetings).Where("account_id = ?", in.AccountID).Scan(ctx); err != nil {
		return nil, err
	}
	return greetings, nil
}

func (r *GreetingRepository) Insert(ctx context.Context, in *model.Greeting) (id int64, err error) {
	res, err := r.conn.NewInsert().Model(in).Exec(ctx)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

package repository

import (
	"context"
	"database/sql"
	"fmt"
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

func (r *GreetingRepository) Update(ctx context.Context, in *model.Greeting) error {

	tx, err := r.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if err := tx.NewSelect().Model(in).WherePK().For("UPDATE").Scan(ctx); err != nil {
		return err
	}

	res, err := tx.NewUpdate().Model(in).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("no target to update: %v", in.ID)
	}
	return nil
}

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"awsomeapp/internal/account/model"

	"github.com/uptrace/bun"
)

type AccountRepository struct {
	conn bun.IDB
}

func NewAccountRepository(conn bun.IDB) *AccountRepository {
	return &AccountRepository{
		conn: conn,
	}
}

func (r *AccountRepository) FindByID(ctx context.Context, in *model.Account) (*model.Account, error) {
	if err := r.conn.NewSelect().Model(in).WherePK().Scan(ctx); err != nil {
		return nil, err
	}
	return in, nil
}

func (r *AccountRepository) Insert(ctx context.Context, in *model.Account) (id int64, err error) {
	res, err := r.conn.NewInsert().Model(in).Exec(ctx)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *AccountRepository) Update(ctx context.Context, in *model.Account) error {

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

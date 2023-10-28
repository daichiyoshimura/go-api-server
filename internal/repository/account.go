package repository

import (
	"awsomeapp/internal/domain/account"
	"awsomeapp/internal/repository/model"
	"context"
	"database/sql"
	"fmt"

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

func (r *AccountRepository) Get(id account.AccountID) (*account.AccountDTO, error) {

	ctx := context.Background()
	ac := &model.Account{
		ID: int64(id),
	}
	if err := r.conn.NewSelect().Model(ac).WherePK().Scan(ctx); err != nil {
		return nil, err
	}
	return ac.DTO(), nil
}

func (r *AccountRepository) Create(in *account.AccountDTO) (*account.AccountDTO, error) {

	ctx := context.Background()
	ac := model.CreateAccountFromDTO(in)

	tx, err := r.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	res, err := tx.NewInsert().Model(ac).Exec(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	ac.ID = id
	return ac.DTO(), nil
}

func (r *AccountRepository) Update(in *account.AccountDTO) (*account.AccountDTO, error) {

	ctx := context.Background()
	ac := model.CreateAccountFromDTO(in)

	tx, err := r.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	if err := tx.NewSelect().Model(ac).WherePK().For("UPDATE").Scan(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	res, err := tx.NewUpdate().Model(ac).WherePK().Exec(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}
	if affected == 0 {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("no target to update: %v", in.ID)
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}
	return ac.DTO(), nil
}

func (r *AccountRepository) Delete(id account.AccountID) error {

	ctx := context.Background()
	ac := &model.Account{
		ID: int64(id),
	}

	tx, err := r.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if err := tx.NewSelect().Model(ac).WherePK().For("UPDATE").Scan(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err := tx.NewDelete().Model(ac).WherePK().Scan(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

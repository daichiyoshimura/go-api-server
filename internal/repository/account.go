package repository

import (
	"context"
	"database/sql"

	"awsomeapp/internal/domain/account"
	"awsomeapp/internal/repository/model"

	"github.com/cockroachdb/errors"
	"github.com/uptrace/bun"
)

const (
	errMsgEmptyTarget string = "command trget does not exist: %w"
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
		return nil, errors.WithStack(err)
	}

	return ac.DTO(), nil
}

func (r *AccountRepository) Create(in *account.AccountDTO) (*account.AccountDTO, error) {
	ctx := context.Background()
	ac := model.CreateAccountFromDTO(in)

	tx, err := r.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := tx.NewInsert().Model(ac).Exec(ctx)
	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.WithStack(err)
		}

		return nil, errors.WithStack(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.WithStack(err)
		}

		return nil, errors.WithStack(err)
	}

	if txerr := tx.Commit(); txerr != nil {
		return nil, errors.WithStack(err)
	}

	ac.ID = id

	return ac.DTO(), nil
}

func (r *AccountRepository) Update(in *account.AccountDTO) (*account.AccountDTO, error) {
	ctx := context.Background()
	ac := model.CreateAccountFromDTO(in)

	tx, err := r.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := tx.NewSelect().Model(ac).WherePK().For("UPDATE").Scan(ctx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.WithStack(err)
		}

		return nil, errors.WithStack(err)
	}

	res, err := tx.NewUpdate().Model(ac).WherePK().Exec(ctx)
	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.WithStack(err)
		}

		return nil, errors.WithStack(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.WithStack(err)
		}

		return nil, errors.WithStack(err)
	}

	if affected == 0 {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.WithStack(err)
		}

		return nil, errors.Newf(errMsgEmptyTarget, in.ID)
	}

	if err := tx.Commit(); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.WithStack(err)
		}

		return nil, errors.WithStack(err)
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
		return errors.WithStack(err)
	}

	if err := tx.NewSelect().Model(ac).WherePK().For("UPDATE").Scan(ctx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return errors.WithStack(err)
		}

		return errors.WithStack(err)
	}

	if err := tx.NewDelete().Model(ac).WherePK().Scan(ctx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return errors.WithStack(err)
		}

		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

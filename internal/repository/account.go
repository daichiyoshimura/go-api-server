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
	errMsgAccountRepository string = "account repository: %w"
	errMsgEmptyResults      string = "query result does not exists: %w"
	errMsgTx                string = "failed to handle transaction: %w"
	errMsgInsert            string = "failed to insert: %w"
	errMsgUpdate            string = "failed to update: %w"
	errMsgEmptyTarget       string = "command trget does not exist: %w"
	errMsgDelete            string = "failed to delete: %w"
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
		return nil, errors.Errorf(errMsgEmptyResults, err)
	}

	return ac.DTO(), nil
}

func (r *AccountRepository) Create(in *account.AccountDTO) (*account.AccountDTO, error) {
	ctx := context.Background()
	ac := model.CreateAccountFromDTO(in)

	tx, err := r.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, errors.Errorf(errMsgTx, err)
	}

	res, err := tx.NewInsert().Model(ac).Exec(ctx)
	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.Errorf(errMsgTx, txerr)
		}

		return nil, errors.Errorf(errMsgInsert, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.Errorf(errMsgTx, txerr)
		}

		return nil, errors.Errorf(errMsgInsert, err)
	}

	if txerr := tx.Commit(); txerr != nil {
		return nil, errors.Errorf(errMsgTx, txerr)
	}

	ac.ID = id

	return ac.DTO(), nil
}

func (r *AccountRepository) Update(in *account.AccountDTO) (*account.AccountDTO, error) {
	ctx := context.Background()
	ac := model.CreateAccountFromDTO(in)

	tx, err := r.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, errors.Errorf(errMsgTx, err)
	}

	if err := tx.NewSelect().Model(ac).WherePK().For("UPDATE").Scan(ctx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.Errorf(errMsgTx, txerr)
		}

		return nil, errors.Errorf(errMsgUpdate, err)
	}

	res, err := tx.NewUpdate().Model(ac).WherePK().Exec(ctx)
	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.Errorf(errMsgTx, txerr)
		}

		return nil, errors.Errorf(errMsgUpdate, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.Errorf(errMsgTx, txerr)
		}

		return nil, errors.Errorf(errMsgUpdate, err)
	}

	if affected == 0 {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.Errorf(errMsgTx, txerr)
		}

		return nil, errors.Errorf(errMsgAccountRepository, errors.Newf(errMsgEmptyTarget, in.ID))
	}

	if err := tx.Commit(); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.Errorf(errMsgTx, txerr)
		}

		return nil, errors.Errorf(errMsgUpdate, err)
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
		return errors.Errorf(errMsgTx, err)
	}

	if err := tx.NewSelect().Model(ac).WherePK().For("UPDATE").Scan(ctx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return errors.Errorf(errMsgTx, txerr)
		}

		return errors.Errorf(errMsgDelete, err)
	}

	if err := tx.NewDelete().Model(ac).WherePK().Scan(ctx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return errors.Errorf(errMsgTx, txerr)
		}

		return errors.Errorf(errMsgDelete, err)
	}

	if err := tx.Commit(); err != nil {
		return errors.Errorf(errMsgTx, err)
	}

	return nil
}

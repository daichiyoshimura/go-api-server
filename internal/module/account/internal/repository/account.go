package repository

import (
	"context"
	"database/sql"

	"awsomeapp/internal/module/account/internal/domain"
	"awsomeapp/internal/module/account/internal/repository/model"

	"github.com/cockroachdb/errors"
	"github.com/uptrace/bun"
)

const (
	sqlUPDATE string = "UPDATE"
)

const (
	errEmptyTarget string = "command target does not exist: %w"
)

type AccountRepository struct {
	db bun.IDB
}

func NewAccountRepository(db bun.IDB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Get(id string) (*domain.AccountDTO, error) {
	ctx := context.Background()
	ac, err := model.CreateAccountFromID(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := r.db.NewSelect().Model(ac).WherePK().Scan(ctx); err != nil {
		return nil, errors.WithStack(err)
	}

	dto, err := ac.DTO()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dto, nil
}

func (r *AccountRepository) Create(udto *domain.AccountUnspecifiedDTO) (*domain.AccountDTO, error) {
	ctx := context.Background()

	ac, err := model.CreateAccountFromUnspecifiedDTO(udto)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if _, err = tx.NewInsert().Model(ac).Exec(ctx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.WithStack(err)
		}

		return nil, errors.WithStack(err)
	}

	if txerr := tx.Commit(); txerr != nil {
		return nil, errors.WithStack(err)
	}

	dto, err := ac.DTO()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dto, nil
}

func (r *AccountRepository) Update(dto *domain.AccountDTO) (*domain.AccountDTO, error) {
	ctx := context.Background()

	ac, err := model.CreateAccountFromDTO(dto)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := tx.NewSelect().Model(ac).WherePK().For(sqlUPDATE).Scan(ctx); err != nil {
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

		return nil, errors.Newf(errEmptyTarget, dto.ID)
	}

	if err := tx.Commit(); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, errors.WithStack(err)
		}

		return nil, errors.WithStack(err)
	}

	retDto, err := ac.DTO()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return retDto, nil
}

func (r *AccountRepository) Delete(id string) error {
	ctx := context.Background()

	ac, err := model.CreateAccountFromID(id)
	if err != nil {
		return errors.WithStack(err)
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return errors.WithStack(err)
	}

	if err := tx.NewSelect().Model(ac).WherePK().For(sqlUPDATE).Scan(ctx); err != nil {
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

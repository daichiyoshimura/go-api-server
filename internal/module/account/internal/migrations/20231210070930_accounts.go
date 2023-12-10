package migrations

import (
	"awsomeapp/internal/module/account/internal/repository/model"
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		if _, err := db.NewCreateTable().Model((*model.Account)(nil)).Exec(ctx); err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		if _, err := db.NewDropTable().Model((*model.Account)(nil)).IfExists().Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}

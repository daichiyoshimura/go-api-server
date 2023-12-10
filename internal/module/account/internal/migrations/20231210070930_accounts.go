package migrations

import (
	"awsomeapp/internal/module/account/internal/repository/model"
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		if _, err := db.NewCreateTable().Model((*model.Account)(nil)).Exec(ctx); err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		if _, err := db.NewDropTable().Model((*model.Account)(nil)).IfExists().Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}

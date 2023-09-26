package dbdriver

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func Establish() (*bun.DB, error) {
	engine, err := sql.Open("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		return nil, err
	}
	return bun.NewDB(engine, mysqldialect.New()), nil
}
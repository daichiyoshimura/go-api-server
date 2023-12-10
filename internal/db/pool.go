package db

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

const (
	driverName string = "mysql"
	net        string = "tcp"
	charset    string = "utf8mb4"
	parseTime  string = "true"
)

type Pool struct{}

func NewPool() *Pool {
	return &Pool{}
}

type IEnv interface {
	Host() string
	User() string
	Password() string
	Name() string
}

func (c *Pool) Establish(env IEnv) (*bun.DB, error) {
	dsn := (&mysql.Config{
		User:      env.User(),
		Passwd:    env.Password(),
		Net:       net,
		Addr:      env.Host(),
		DBName:    env.Name(),
		ParseTime: true,
	}).FormatDSN()

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, errors.WithStack(err)
	}

	return bun.NewDB(db, mysqldialect.New()), nil
}

package db

import (
	"context"
	"database/sql"
	"strings"

	"github.com/cockroachdb/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

const (
	driverName string = "mysql"
	protocol   string = "tcp"
	charset    string = "utf8mb4"
	parseTime  string = "true"
)

const (
	errMsgPool string = "failed to establish db Pool: %w"
)

type Pool struct{}

func NewPool() *Pool {
	return &Pool{}
}

type IEnv interface {
	Host() string
	User() string
	Password() string
	Instance() string
}

func (c *Pool) Establish(env IEnv) (*bun.DB, error) {
	source, err := c.dataSourceName(env)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db, err := sql.Open(driverName, source)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, errors.WithStack(err)
	}

	return bun.NewDB(db, mysqldialect.New()), nil
}

// TODO: use mysql.config.
func (c *Pool) dataSourceName(env IEnv) (string, error) {
	parts := []string{
		env.User(),
		":",
		env.Password(),
		"@",
		protocol,
		"(",
		env.Host(),
		")/",
		env.Instance(),
	}

	options := map[string]string{
		"charset":   charset,
		"parseTime": parseTime,
	}
	c.appendOptions(&parts, &options)

	sb := &strings.Builder{}
	for _, part := range parts {
		if _, err := sb.WriteString(part); err != nil {
			return "", errors.WithStack(err)
		}
	}

	return sb.String(), nil
}

func (c *Pool) appendOptions(parts *[]string, options *map[string]string) {
	partsOriginLength := len(*parts)

	var partsLength int
	for k, v := range *options {
		partsLength = len(*parts)
		if partsOriginLength == partsLength {
			*parts = append(*parts, "?")
		} else {
			*parts = append(*parts, "&")
		}

		*parts = append(*parts, k)
		*parts = append(*parts, "=")
		*parts = append(*parts, v)
	}
}

package db

import (
	"database/sql"
	"strings"

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

type Connection struct{}

func NewConnection() *Connection {
	return &Connection{}
}

type IEnv interface {
	Host() string
	User() string
	Password() string
	Instance() string
}

func (c *Connection) Establish(env IEnv) (*bun.DB, error) {
	source, err := c.dataSourceName(env)
	if err != nil {
		return nil, err
	}

	conn, err := sql.Open(driverName, source)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return bun.NewDB(conn, mysqldialect.New()), nil
}

func (c *Connection) dataSourceName(env IEnv) (string, error) {
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
			return "", err
		}
	}

	return sb.String(), nil
}

func (c *Connection) appendOptions(parts *[]string, options *map[string]string) {
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

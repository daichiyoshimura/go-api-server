package db

import (
	"database/sql"
	"strings"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName string = "mysql"
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

	sb := &strings.Builder{}
	if _, err := sb.WriteString(env.User()); err != nil {
		return "", err
	}
	if _, err := sb.WriteString(":"); err != nil {
		return "", err
	}
	if _, err := sb.WriteString(env.Password()); err != nil {
		return "", err
	}
	if _, err := sb.WriteString("@tcp("); err != nil {
		return "", err
	}
	if _, err := sb.WriteString(env.Host()); err != nil {
		return "", err
	}
	if _, err := sb.WriteString(")/"); err != nil {
		return "", err
	}
	if _, err := sb.WriteString(env.Instance()); err != nil {
		return "", err
	}
	if _, err := sb.WriteString("?charset=utf8mb4&parseTime=true"); err != nil {
		return "", err
	}
	return sb.String(), nil
}

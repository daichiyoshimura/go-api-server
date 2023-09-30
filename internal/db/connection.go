package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {}

func NewConnection() *Connection{
	return &Connection{}
}

func (c *Connection) Establish() (*bun.DB, error) {
	conn, err := sql.Open("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return bun.NewDB(conn, mysqldialect.New()), nil
}
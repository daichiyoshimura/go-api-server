package dbdriver

import "github.com/uptrace/bun"

type Greetings struct {
	driver *bun.IDB
}

func NewGreetings(driver *bun.IDB) *Greetings {
	return &Greetings{
		driver: driver,
	}
}
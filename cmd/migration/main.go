package main

import "github.com/uptrace/bun/migrate"

func main() {
	if err := migrate.NewMigrations().DiscoverCaller(); err != nil {
		panic(err)
	}
}

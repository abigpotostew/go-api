package migration

import (
	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations()

//var sqlMigrations embed.FS

func init() {
	if err := Migrations.DiscoverCaller(); err != nil {
		panic(err)
	}
}

package storage

import (
	"embed"
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type Migration struct {
}

type MigrationI interface {
	Up()
	Down()
}

//go:embed migrations/*.sql
var fs embed.FS

func (m Migration) Up() {
	db := DB{}

	conn, err := db.connect()
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	driver, err := mysql_migrate.WithInstance(conn, &mysql_migrate.Config{})

	if err != nil {
		panic(err.Error())
	}

	iofsPath, err := iofs.New(fs, "migrations")
	if err != nil {
		panic(err.Error())
	}
	migration, err := migrate.NewWithInstance(
		"iofs",
		iofsPath,
		"mysql",
		driver,
	)

	if err != nil {
		panic(err.Error())
	}
	err = migration.Up()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Migrations up!")
}

func (m Migration) Down() {
}

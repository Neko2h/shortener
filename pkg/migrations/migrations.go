package migrations

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(migrationPath string, Db string) error {

	if migrationPath == "" {
		return errors.New("no migration path was provided")
	}
	if Db == "" {
		return errors.New("no DB url was provided")
	}
	m, err := migrate.New(
		migrationPath,
		Db,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

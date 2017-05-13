package database

import (
	"errors"
	"path"
	"runtime"
	"strings"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres" // for postgres compatibility
	_ "github.com/mattes/migrate/source/file"       // file system migrations
)

// NewMigration returns a new migration instance
func NewMigration(conn string) (*migrate.Migrate, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return &migrate.Migrate{}, errors.New("No caller information")
	}

	migrationsPath := strings.Join([]string{"file:/", path.Dir(filename), "migrations"}, "/")
	m, err := migrate.New(migrationsPath, getConnectionString(conn))
	if err != nil {
		panic(err)
	}

	return m, nil
}

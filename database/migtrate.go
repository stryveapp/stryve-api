package database

import (
	"path"
	"runtime"
	"strings"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

// NewMigration returns a new migration instance
func NewMigration(conn string) *migrate.Migrate {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	migrationsPath := strings.Join([]string{"file:/", path.Dir(filename), "migrations"}, "/")
	m, err := migrate.New(migrationsPath, getConnectionString(conn))
	if err != nil {
		panic(err)
	}

	return m
}

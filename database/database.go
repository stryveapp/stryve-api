package database

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/stryveapp/stryve-api/config"
)

// DB is the database layer
type DB struct {
}

// Open opens the DB connection
func Open() *pg.DB {
	var db *DB
	return db.OpenConnection(config.Env)
}

// Close closes the provided database connection
func (db *DB) Close(conn *pg.DB) {
	if err := conn.Close(); err != nil {
		panic(err)
	}
}

// GetConnection returns the info required to establish a DB connection
func (db *DB) GetConnection(connEnv string) *pg.DB {
	connInfo := config.DB[connEnv]

	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", connInfo.Host, connInfo.Port),
		Database: connInfo.Name,
		User:     connInfo.Username,
		Password: connInfo.Password,
	})
}

// OpenConnection opens and returns a new DB connection
// based on the environment passed in
func (db *DB) OpenConnection(connEnv string) *pg.DB {
	return db.GetConnection(connEnv)
}

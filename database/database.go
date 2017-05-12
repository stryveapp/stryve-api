package database

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/stryveapp/stryve-api/config"
)

const (
	DuplicateKeyViolationError = "23505"
)

// NewConnection opens and retruns the specified DB connection
func NewConnection(connection ...string) *pg.DB {
	connInfo := getConnectionInfo(connection...)

	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", connInfo.Host, connInfo.Port),
		Database: connInfo.Name,
		User:     connInfo.Username,
		Password: connInfo.Password,
	})
}

func getConnectionString(connection ...string) string {
	connInfo := getConnectionInfo(connection...)

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		connInfo.Username,
		connInfo.Password,
		connInfo.Host,
		connInfo.Port,
		connInfo.Name,
		connInfo.SSLMode,
	)
}

func getConnectionInfo(connection ...string) config.DatabaseConfig {
	var conn string

	if len(connection) == 1 {
		connectionTypes := []string{"dev", "test", "prod"}
		for _, connType := range connectionTypes {
			if connection[0] == connType {
				conn = connType
				break
			}
		}
	}

	if conn == "" {
		conn = config.Env
	}

	return config.DB[conn]
}

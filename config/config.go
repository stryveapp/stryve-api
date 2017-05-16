package config

import (
	"errors"
	"path"
	"runtime"
	"strings"

	"github.com/BurntSushi/toml"
)

// Config get populated with the configurations
// from the config.toml file
var Config config

type config struct {
	// Env is the application set environment
	Env string
	// Port is the HTTP port to run the server on
	Port int
	// Debug enables/disable application verbosity
	Debug bool
	// LogLevel is the level at which logs will be produced
	LogLevel string `toml:"log_level"`
	// LogFormat is the format that logs will be written in
	LogFormat string `toml:"log_format"`
	// LogPath is the absolute file path to the directory where logs will be stored
	LogPath string `toml:"log_path"`
	// LogFilename is the name given to the log file
	LogFilename string `toml:"log_filename"`
	// DB is a list of available database connections
	Databases map[string]DatabaseConfig
}

// DatabaseConfig is the configuration set for
// a PostgresSQL connection
type DatabaseConfig struct {
	Host     string
	Port     int
	SSLMode  string `toml:"ssl_mode"`
	Name     string
	Username string
	Password string
}

// LoadDefaultConfig sets the servers default configuration set
func LoadDefaultConfig() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("No caller information")
	}

	path := strings.Join([]string{path.Dir(filename), "config.toml"}, "/")
	if _, err := toml.DecodeFile(path, &Config); err != nil {
		return err
	}

	return nil
}

package config

import (
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/stryveapp/stryve-api/util"
)

var (
	// Env is the application set environment
	Env string

	// Debug enables/disable application verbosity
	Debug bool

	// Port is the HTTP port to run the server on
	Port int

	// DB is a list of availble datbase connections
	DB map[string]databaseConfig
)

type config struct {
	Env       string
	Port      int
	Debug     bool
	Databases map[string]databaseConfig
}

type databaseConfig struct {
	Host     string
	Port     int
	SSLMode  string `toml:"ssl_mode"`
	Name     string
	Username string
	Password string
}

// SetDefaultConfig set the servers default configuration set
func SetDefaultConfig() {
	var conf config

	path := strings.Join([]string{util.GetCWD(), "config", "config.toml"}, "/")
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		panic(err)
	}

	Env = conf.Env
	Debug = conf.Debug
	Port = conf.Port
	DB = conf.Databases
}

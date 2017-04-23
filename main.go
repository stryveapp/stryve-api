package main

import (
	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/server"
	"github.com/stryveapp/stryve-api/validate"
)

func init() {
	config.SetDefaultConfig()
	validate.RegisterCustomValidators()
}

func main() {
	svr := server.New()
	server.StartServer(svr)
}

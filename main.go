package main

import (
	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/server"
)

func init() {
	config.SetDefaultConfig()
}

func main() {
	e := server.New()
	server.StartServer(e)
}

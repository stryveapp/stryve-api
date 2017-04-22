package main

import (
	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/server"
)

func init() {
	config.SetDefaultConfig()
}

func main() {
	svr := server.New()
	server.StartServer(svr)
}

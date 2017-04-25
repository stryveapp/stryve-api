package main

import (
	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/server"
)

func init() {
	err := config.SetDefaultConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	e := server.New()
	server.StartServer(e)
}

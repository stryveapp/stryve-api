package main

import (
	"github.com/stryveapp/stryve-api/router"
	"github.com/stryveapp/stryve-api/server"
)

func main() {
	svr := server.New()
	router.RegisterRoutes(svr)
	server.StartServer(svr)
}

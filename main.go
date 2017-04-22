package main

import (
	"github.com/stryveapp/stryve-api/server"
)

func main() {
	svr := server.New()
	server.StartServer(svr)
}

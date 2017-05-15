package main

import (
	"log"
	"os"

	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/server"
)

func init() {
	err := config.LoadDefaultConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	e := server.New()
	server.StartServer(e)
}

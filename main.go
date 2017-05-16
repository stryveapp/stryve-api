package main

import (
	"flag"
	"log"

	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/logger"
	"github.com/stryveapp/stryve-api/server"
)

func init() {
	err := config.LoadDefaultConfig()
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	flag.StringVar(&config.Config.Env, "env", config.Config.Env, "the application running environment")
	flag.BoolVar(&config.Config.Debug, "debug", config.Config.Debug, "enable or disable debug mode")
	flag.StringVar(&config.Config.LogLevel, "log-level", config.Config.LogLevel, "the level at which logs will be produced")
	flag.StringVar(&config.Config.LogFormat, "log-format", config.Config.LogFormat, "the format that logs will be written in")
	flag.StringVar(&config.Config.LogPath, "log-path", config.Config.LogPath, "the absolute file path to the logs directory")
	flag.StringVar(&config.Config.LogFilename, "log-filename", config.Config.LogFilename, "is the name given to the log file")
	flag.Parse()

	logger.InitLogger()
}

func main() {
	e := server.New()
	server.StartServer(e)
}

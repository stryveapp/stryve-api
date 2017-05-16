package logger

import (
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/stryveapp/stryve-api/config"
)

// Log instantiates a new logger
var Log = logrus.New()

// InitLogger intantiates a new logger
func InitLogger() {
	Log.Level = getLogLevel(config.Config.LogLevel)
	Log.Formatter = getLogFormat(config.Config.LogFormat)

	var err error
	Log.Out, err = getLogLocation(config.Config.LogPath)
	if err != nil {
		Log.Info("Failed to log to file, falling back to default stderr")
	}
}

func getLogLocation(path string) (*os.File, error) {
	var filepath string

	if path == "" {
		cwd, _ := os.Getwd()
		filepath = strings.Join([]string{cwd, "logs", config.Config.LogFilename}, "/")
	} else {
		filepath = strings.Join([]string{path, config.Config.LogFilename}, "/")
	}

	err := os.MkdirAll(
		strings.TrimSuffix(filepath, strings.Join([]string{"/", config.Config.LogFilename}, "")),
		os.ModePerm,
	)
	if err != nil {
		return &os.File{}, err
	}

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return &os.File{}, err
	}

	return file, nil
}

func getLogFormat(format string) logrus.Formatter {
	defaultFormatter := &logrus.TextFormatter{
		DisableColors:  true,
		DisableSorting: true,
	}

	switch strings.ToLower(format) {
	case "text":
		return defaultFormatter
	case "json":
		return &logrus.JSONFormatter{}
	default:
		return defaultFormatter
	}
}

func getLogLevel(level string) logrus.Level {
	defaultLogLevel := logrus.InfoLevel

	switch strings.ToLower(level) {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return defaultLogLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return defaultLogLevel
	}
}

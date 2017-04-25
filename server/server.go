package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/database"
	"github.com/stryveapp/stryve-api/middleware"
	"github.com/stryveapp/stryve-api/router"
)

// New creates a new instance of App, which
// is an extension of Echo
func New() (e *echo.Echo) {
	e = echo.New()
	e.Debug = config.Debug
	e.Logger.SetLevel(log.INFO)
	router.RegisterRoutes(e)
	middleware.RegisterMiddleware(e)

	return e
}

// StartServer starts the API server
func StartServer(e *echo.Echo) {
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", config.Port)); err != nil {
			e.Logger.Info("Unable to start server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	StopServer(e)
}

// StopServer stops the API server gracefully
func StopServer(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := database.NewConnection()
	defer db.Close() // close any and all DB connections

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

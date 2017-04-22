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
	"github.com/stryveapp/stryve-api/middleware"
	"github.com/stryveapp/stryve-api/router"
)

// New creates a new instance of App, which
// is an extension of Echo
func New() (svr *echo.Echo) {
	svr = echo.New()
	svr.Debug = config.Debug
	router.RegisterRoutes(svr)
	middleware.RegisterMiddleware(svr)
	svr.Logger.SetLevel(log.INFO)

	return svr
}

// StartServer starts the API server
func StartServer(svr *echo.Echo) {
	go func() {
		if err := svr.Start(fmt.Sprintf(":%d", config.Port)); err != nil {
			svr.Logger.Info("Unable to start server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	StopServer(svr)
}

// StopServer stops the API server gracefully
func StopServer(svr *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		svr.Logger.Fatal(err)
	}
}

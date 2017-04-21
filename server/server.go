package server

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/stryveapp/stryve-api/middleware"
)

// CustomContext is a custom context that extends Echo's
// base Context implementation
type CustomContext struct {
	echo.Context
}

// New creates a new instance of App, which
// is an extension of Echo
func New() (svr *echo.Echo) {
	svr = echo.New()
	svr.Logger.SetLevel(log.INFO)

	// Middleware
	svr.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return h(cc)
		}
	})
	svr.Use(middleware.Logger())
	svr.Use(middleware.Recover())

	return svr
}

// StartServer starts the API server
func StartServer(svr *echo.Echo) {
	go func() {
		if err := svr.Start(":3000"); err != nil {
			svr.Logger.Info("shutting down the server")
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

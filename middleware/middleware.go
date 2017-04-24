package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RegisterMiddleware registers all required middleware
func RegisterMiddleware(svr *echo.Echo) {
	svr.Pre(middleware.RemoveTrailingSlash())
	svr.Use(middleware.Logger())
	svr.Use(middleware.Recover())
}

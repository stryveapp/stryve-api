package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RegisterMiddleware registers all required middleware
func RegisterMiddleware(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

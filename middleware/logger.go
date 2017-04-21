package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Logger is Echo's base logger implementation
func Logger() echo.MiddlewareFunc {
	return middleware.Logger()
}

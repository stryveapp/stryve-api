package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Recover is Echo's base recover implementation
func Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}

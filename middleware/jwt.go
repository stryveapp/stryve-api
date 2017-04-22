package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// JWT is Echo's base JWT implementation
func JWT(secret []byte) echo.MiddlewareFunc {
	return middleware.JWT(secret)
}

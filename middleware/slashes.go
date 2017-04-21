package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RemoveTrailingSlash is Echo's base remove slashes implementation
func RemoveTrailingSlash() echo.MiddlewareFunc {
	return middleware.RemoveTrailingSlash()
}

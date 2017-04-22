package middleware

import "github.com/labstack/echo"

// CustomContext is a custom context that extends Echo's
// base Context implementation
type CustomContext struct {
	echo.Context
}

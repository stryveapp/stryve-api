package router

import (
	"net/http"

	"github.com/labstack/echo"
)

// RegisterRoutes registers all API routes for the app
func RegisterRoutes(svr *echo.Echo) {
	svr.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
}

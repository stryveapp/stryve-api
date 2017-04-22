package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Login is the login endpoint handler
func Login(c echo.Context) error {
	return c.String(http.StatusOK, "LOGIN")
}

// Logout is the logout endpoint handler
func Logout(c echo.Context) error {
	return c.String(http.StatusOK, "LOGOUT")
}

// Register is the logout endpoint handler
func Register(c echo.Context) error {
	return c.String(http.StatusOK, "LOGOUT")
}

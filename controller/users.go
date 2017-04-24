package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetUsers list all communities
func (h *Handler) GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "LISTING USERS")
}

// GetUser returns a single user
func (h *Handler) GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "GETTING USERS")
}

// CreateUser creates a user
func (h *Handler) CreateUser(c echo.Context) error {
	return c.String(http.StatusOK, "CREATING USERS")
}

// UpdateUser updates a user
func (h *Handler) UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "UPDATING USERS")
}

// DeleteUser soft deletes a user
func (h *Handler) DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "DELETING USERS")
}

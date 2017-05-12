package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/stryveapp/stryve-api/model"
	"github.com/stryveapp/stryve-api/repository"
)

// Login is the login endpoint handler
func (h *Handler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "LOGIN")
}

// Logout is the logout endpoint handler
func (h *Handler) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "LOGOUT")
}

// Register is the logout endpoint handler
func (h *Handler) Register(c echo.Context) error {
	user := new(model.User)
	c.Bind(&user)

	repo := repository.UsersRepo{
		DB: h.DB,
	}

	errs := repo.CreateNewUser(user)
	if len(errs) > 0 {
		return c.JSON(http.StatusBadRequest, NewHTTPError(errs))
	}

	return c.JSON(http.StatusCreated, NewHTTPSuccess(nil))
}

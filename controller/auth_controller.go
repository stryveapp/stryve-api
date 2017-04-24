package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/stryveapp/stryve-api/model"
	"github.com/stryveapp/stryve-api/request"
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
	v, errs := request.ValidateRegisterRequest(h.DB, c)
	if len(errs) > 0 {
		return c.JSON(http.StatusBadRequest, NewHTTPError(errs))
	}

	params := v.(request.RegisterRequest)

	user := new(model.User)
	user.Username = params.Username
	user.Email = params.Email
	user.Password = params.Password
	user.Save(h.DB)

	return c.JSON(http.StatusCreated, NewHTTPSuccess(nil))
}

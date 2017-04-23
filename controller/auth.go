package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/stryveapp/stryve-api/model"
	"github.com/stryveapp/stryve-api/repository"
	"github.com/stryveapp/stryve-api/request"
)

type loginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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
	v, errs := request.ValidateRequest("authRegister", c)
	if len(errs) > 0 {
		return c.JSON(http.StatusOK, NewHTTPError(errs))
	}

	now := time.Now()
	params := v.(request.RegisterRequest)

	user := new(model.User)
	user.Username = params.Username
	user.Email = params.Email
	user.Password = repository.GeneratePasswordHash(params.Password)
	user.VerificationToken = repository.GenerateEmailVerificationToken(60)
	user.CreatedAt = now
	user.UpdatedAt = now
	user.Save()

	return c.JSON(http.StatusCreated, NewHTTPSuccess(nil))
}

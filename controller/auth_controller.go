package controller

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/stryveapp/stryve-api/model"
	"github.com/stryveapp/stryve-api/validator"
)

// Login is the login endpoint handler
func (h *Handler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "LOGIN")
}

// Logout is the logout endpoint handler
func (h *Handler) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "LOGOUT")
}

type registerRequest struct {
	Username    string
	DisplayName string
	Email       string
	Password    string
}

// Register is the logout endpoint handler
func (h *Handler) Register(c echo.Context) error {
	var params *registerRequest
	c.Bind(&params)

	params.DisplayName = params.Username
	params.Username = strings.ToLower(params.Username)

	errs := validateRegisterRequest(h.DB, params)
	if len(errs) > 0 {
		return c.JSON(http.StatusBadRequest, NewHTTPError(errs))
	}

	user := new(model.User)
	user.Username = params.Username
	user.DisplayName = params.DisplayName
	user.Email = params.Email
	user.Password = params.Password
	user.Save(h.DB)

	return c.JSON(http.StatusCreated, NewHTTPSuccess(nil))
}

// ValidateRegisterRequest validates the /auth/register POST request
func validateRegisterRequest(db *pg.DB, params *registerRequest) []string {
	v := &validator.Validator{DB: db}

	if params.Email == "" || params.Username == "" || params.Password == "" {
		return append(v.Errors, "All feilds are required")
	}

	// VALIDATE USERNAME
	if !v.IsValidUsername(params.Username) {
		v.Errors = append(v.Errors, "Username can only contain alpanumeric and underscore characters, and must be between 3 and 25 characters long")
	}
	if !v.IsUniqueUsername(params.Username) {
		v.Errors = append(v.Errors, "Username is already taken")
	}

	// VALIDATE EMAIL
	if !v.IsValidEmail(params.Email) {
		v.Errors = append(v.Errors, "Invalid email address")
	}
	if !v.IsUniqueEmail(params.Email) {
		v.Errors = append(v.Errors, "Email address is already registered")
	}

	// VALIDATE PASSWORD
	if !v.IsValidPassword(params.Password) {
		v.Errors = append(v.Errors, "Password must contain both numbers and letters, and be a minimum of 8 characters long")
	}

	return v.Errors
}

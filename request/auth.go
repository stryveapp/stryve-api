package request

import (
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/stryveapp/stryve-api/validator"
)

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

// ValidateRegisterRequest validates the /auth/register POST request
func ValidateRegisterRequest(db *pg.DB, c echo.Context) (interface{}, []string) {
	var req RegisterRequest
	c.Bind(&req)

	v := &validator.Validator{DB: db}

	if req.Email == "" || req.Username == "" || req.Password == "" {
		v.Errors = append(v.Errors, "All feilds are required")
	}

	// VALIDATE USERNAME
	if !v.IsValidStringLength(req.Username, 8, 25) {
		v.Errors = append(v.Errors, "Username must be between 3 and 25 characters long")
	}
	if !v.IsValidUsername(req.Username) {
		v.Errors = append(v.Errors, "Username can only contain alpanumeric and underscore characters")
	}
	if !v.IsUniqueUsername(req.Username) {
		v.Errors = append(v.Errors, "Username is already taken")
	}

	// VALIDATE EMAIL
	if !v.IsValidEmail(req.Email) {
		v.Errors = append(v.Errors, "Invalid email address")
	}
	if !v.IsUniqueEmail(req.Email) {
		v.Errors = append(v.Errors, "Email is already registered")
	}

	// VALIDATE PASSWORD
	if !v.IsValidPassword(req.Password) {
		v.Errors = append(v.Errors, "Password must contain both numbers and letters, and be a minimum of 8 characters long")
	}

	return req, v.Errors
}

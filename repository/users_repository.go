package repository

import (
	"strings"

	"github.com/go-pg/pg"
	"github.com/stryveapp/stryve-api/database"
	"github.com/stryveapp/stryve-api/model"
	"github.com/stryveapp/stryve-api/validator"
)

type UsersRepo struct {
	DB *pg.DB
}

// CreateNewUser validates and creates a new user
func (up *UsersRepo) CreateNewUser(user *model.User) []string {
	v := &validator.Validator{}

	if user.Email == "" || user.Username == "" || user.Password == "" {
		return append(v.Errors, "All feilds are required")
	}

	// VALIDATE USERNAME
	if !v.IsValidUsername(user.Username) {
		v.Errors = append(v.Errors, "Username can only contain alpanumeric and underscore characters, and must be between 3 and 25 characters long")
	}

	// VALIDATE EMAIL
	if !v.IsValidEmail(user.Email) {
		v.Errors = append(v.Errors, "Invalid email address")
	}

	// VALIDATE PASSWORD
	if !v.IsValidPassword(user.Password) {
		v.Errors = append(v.Errors, "Password must contain both numbers and letters, and be a minimum of 8 characters long")
	}

	// ATTEMPT TO CREATE NEW USER
	if err := user.Insert(up.DB); err != nil {
		if strings.Contains(err.Error(), database.DuplicateKeyViolationError) {
			v.Errors = append(v.Errors, "Your email address and/or username are already registered")
		} else {
			v.Errors = append(v.Errors, "An unknown error occurred")
		}
	}

	return v.Errors
}

// CountUsersBy returns number of users that match the
// column and value provided
func (up *UsersRepo) CountUsersBy(column, value string) int {
	count, _ := up.DB.Model(&model.User{}).
		Where("? = ?", column, value).
		Count()

	return count
}

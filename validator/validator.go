package validator

import (
	"regexp"
	"unicode"

	"github.com/go-pg/pg"
	"github.com/stryveapp/stryve-api/model"
	"github.com/stryveapp/stryve-api/util"
)

type Validator struct {
	DB     *pg.DB
	Errors []string
}

// IsValidStringLength return true if the provided string's length
// is within the min and max values
func (v *Validator) IsValidStringLength(str string, min, max int) bool {
	return len(str) >= min && len(str) <= max
}

// IsValidUsername validates if a provided string is
// a valid username
func (v *Validator) IsValidUsername(str string) bool {
	reg := regexp.MustCompile(util.UsernameRegex)

	return reg.MatchString(str)
}

// IsValidEmail validates if a provided string is
// a valid email
func (v *Validator) IsValidEmail(str string) bool {
	reg := regexp.MustCompile(util.EmailRegex)

	return reg.MatchString(str)
}

// IsValidPassword validates if a provided string is
// a valid password
func (v *Validator) IsValidPassword(str string) bool {
	if len(str) < 8 {
		return false
	}

	hasNumbers := false
	hasLetters := false

	for _, char := range str {
		switch {
		case unicode.IsNumber(char):
			hasNumbers = true
		case unicode.IsLetter(char):
			hasLetters = true
		default:
		}
	}

	return hasLetters && hasNumbers
}

// IsUniqueEmail returns true is the provided string is
// a unique email address in the users table
func (v *Validator) IsUniqueEmail(str string) bool {
	count, _ := v.DB.Model(&model.User{}).
		Where("email = ?", str).
		Count()

	return count < 1
}

// IsUniqueUsername returns true is the provided string is
// a unique username in the users table
func (v *Validator) IsUniqueUsername(str string) bool {
	count, _ := v.DB.Model(&model.User{}).
		Where("username = ?", str).
		Count()

	return count < 1
}

package validator

import (
	"regexp"
	"unicode"

	"github.com/stryveapp/stryve-api/util"
)

type Validator struct {
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

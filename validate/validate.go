package validate

import (
	"regexp"
	"unicode"

	"github.com/asaskevich/govalidator"
	"github.com/stryveapp/stryve-api/database"
	"github.com/stryveapp/stryve-api/model"
	"github.com/stryveapp/stryve-api/util"
)

func RegisterCustomValidators() {
	govalidator.TagMap["username"] = govalidator.Validator(Username)
	govalidator.TagMap["password"] = govalidator.Validator(Password)
	govalidator.TagMap["unique_email"] = govalidator.Validator(UniqueUserEmail)
	govalidator.TagMap["unique_username"] = govalidator.Validator(UniqueUserUsername)
}

// Username validates if a provided string is
// a valid username
func Username(str string) bool {
	matchme := regexp.MustCompile(util.UsernameRegex)

	return matchme.MatchString(str)
}

// Password validates if a provided string is
// a valid password
func Password(str string) bool {
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

// UniqueUserEmail returns true is the provided string is
// a unique email address in the users table
func UniqueUserEmail(str string) bool {
	db := database.Open()
	defer db.Close()

	count, _ := db.Model(&model.User{}).
		Where("email = ?", str).
		Count()

	return count < 1
}

// UniqueUserUsername returns true is the provided string is
// a unique username in the users table
func UniqueUserUsername(str string) bool {
	db := database.Open()
	defer db.Close()

	count, _ := db.Model(&model.User{}).
		Where("username = ?", str).
		Count()

	return count < 1
}

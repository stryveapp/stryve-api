package repository

import (
	"github.com/stryveapp/stryve-api/database"
	"github.com/stryveapp/stryve-api/model"
	"github.com/stryveapp/stryve-api/util"
	"golang.org/x/crypto/bcrypt"
)

//GeneratePasswordHash generates and retruns a
// new password hash
func GeneratePasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash)
}

// GenerateEmailVerificationToken generates and retruns a
// new unique email verification token
func GenerateEmailVerificationToken(length int) string {
	var token string

	for {
		token = util.GenerateRandomStrig(60, true, true, true, false)

		db := database.Open()
		defer db.Close()

		count, _ := db.Model(model.User{}).
			Where("verification_token = ?", token).
			Count()

		if count < 1 {
			break
		}
	}

	return token
}

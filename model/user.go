package model

import (
	"strings"
	"time"

	"github.com/go-pg/pg/orm"
	"github.com/stryveapp/stryve-api/util"
	"golang.org/x/crypto/bcrypt"
)

// User is the user datasource skeleton
type User struct {
	PrimaryID
	Username          string `json:"username"`
	DisplayName       string `json:"display_name"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Avatar            string `json:"avatar"`
	VerificationToken string `json:"verification_token"`
	CanLogIn          bool   `json:"can_log_in"`
	CommonDates
}

/*
 * HOOKS
 */

// BeforeInsert is the User model before insert hook
func (user *User) BeforeInsert(db orm.DB) error {
	now := time.Now()

	user.DisplayName = user.Username
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)
	user.Password = generatePasswordHash(user.Password)
	user.VerificationToken = generateEmailVerificationToken(db, 60)
	user.CreatedAt = now
	user.UpdatedAt = now

	return nil
}

/*
 * METHODS
 */

// Insert inserts a new user
func (user *User) Insert(db orm.DB) error {
	err := db.Insert(&user)
	if err != nil {
		return err
	}

	return nil
}

/*
 * HELPERS
 */

func generateEmailVerificationToken(db orm.DB, length int) string {
	var token string
	for {
		token = util.GenerateRandomStrig(60, true, true, true, false)

		count, _ := db.Model(User{}).
			Where("verification_token = ?", token).
			Count()

		if count < 1 {
			break
		}
	}

	return token
}

func generatePasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash)
}

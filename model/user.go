package model

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/stryveapp/stryve-api/util"
	"golang.org/x/crypto/bcrypt"
)

// User is the user datasource skeleton
type User struct {
	PrimaryID
	Username          string `json:"username"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Avatar            string `json:"avatar"`
	VerificationToken string `json:"verification_token"`
	CommonDates
}

func (user *User) Save(db *pg.DB) {
	err := db.Insert(&user)
	if err != nil {
		panic(err)
	}
}

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

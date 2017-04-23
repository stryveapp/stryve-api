package model

import (
	"github.com/stryveapp/stryve-api/database"
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

func (user *User) Save() {
	db := database.Open()
	defer db.Close()

	err := db.Insert(&user)
	if err != nil {
		panic(err)
	}
}

package model

import (
	"strings"
	"time"

	"github.com/go-pg/pg/orm"
)

// BeforeInsert modifies and instantiates User fields
// before inserting into the table
func (user *User) BeforeInsert(db orm.DB) error {
	now := time.Now()

	user.Email = strings.ToLower(user.Email)
	user.Password = generatePasswordHash(user.Password)
	user.VerificationToken = generateEmailVerificationToken(db, 60)
	user.CreatedAt = now
	user.UpdatedAt = now

	return nil
}

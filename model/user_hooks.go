package model

import (
	"time"

	"github.com/go-pg/pg/orm"
)

func (user *User) BeforeInsert(db orm.DB) error {
	now := time.Now()

	user.Password = generatePasswordHash(user.Password)
	user.VerificationToken = generateEmailVerificationToken(db, 60)
	user.CreatedAt = now
	user.UpdatedAt = now

	return nil
}

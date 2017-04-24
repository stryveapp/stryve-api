package util

import (
	"math/rand"
)

const (
	// UsernameRegex is the regex pattern for determining valid usernames
	UsernameRegex = "^[a-z0-9_]{3,25}$"
	// EmailRegex is the regex pattern for determining valid email addresses
	EmailRegex = `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`
)

// GenerateRandomStrig generates and returns a random string
func GenerateRandomStrig(length int, numbers, lcase, ucase, specials bool) string {
	var chars []byte
	numChars := "0123456789"
	lcaseChars := "abcdefghijklmnopqrstuvwxyz"
	ucaseChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialChars := "!@#$%&*?"

	if numbers {
		chars = append(chars, numChars...)
	}
	if lcase {
		chars = append(chars, lcaseChars...)
	}
	if ucase {
		chars = append(chars, ucaseChars...)
	}
	if specials {
		chars = append(chars, specialChars...)
	}

	random := make([]byte, length)
	for i := 0; i < length; i++ {
		random[i] = chars[rand.Intn(len(chars))]
	}

	return string(random)
}

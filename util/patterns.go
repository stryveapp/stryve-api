package util

const (
	// UsernameRegex is the regex pattern for determining valid usernames
	UsernameRegex = "^[a-zA-Z0-9_]{3,25}$"
	// EmailRegex is the regex pattern for determining valid email addresses
	EmailRegex = `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`
)

package model

// User is the user datasource skelton
type User struct {
	PrimaryID
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	UserStatus Status `json:"user_status"`
	CommonDates
}

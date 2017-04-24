package validator_test

import (
	"testing"

	"github.com/mattes/migrate"
	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/database"
	"github.com/stryveapp/stryve-api/validator"
)

var (
	v *validator.Validator
	m *migrate.Migrate
)

func setUp() {
	config.SetDefaultConfig()
	m = database.NewMigration("test")
	m.Up()

	v = &validator.Validator{
		DB: database.NewConnection("test"),
	}
}

func tearDown() {
	defer v.DB.Close()
	m.Down()
}

func TestIsValidStringLength(t *testing.T) {
	testCases := []struct {
		str     string
		min     int
		max     int
		isValid bool
	}{
		{"aa", 3, 4, false},
		{"aaa", 3, 4, true},
	}

	for _, tc := range testCases {
		ok := v.IsValidStringLength(tc.str, tc.min, tc.max)
		if ok != tc.isValid {
			t.Fatalf(`String "%s" expects to be >= %d and <= %d in length, but is not`, tc.str, tc.min, tc.max)
		}
	}
}

func TestIsValidUsersame(t *testing.T) {
	testCases := []struct {
		username string
		isValid  bool
	}{
		{"aa", false},                         // too short, min 3 characters required
		{"sa1d31asd321a32s1d3sssssss", false}, // too long, max 25 characters required
		{"inv4lid-name", false},               // no dash characters allowed
		{"us3rn4me!", false},                  // no special characters allowed
		{"12345678", true},
		{"0_____0", true},
		{"_______", true},
	}

	for _, tc := range testCases {
		ok := v.IsValidUsername(tc.username)
		if ok != tc.isValid {
			t.Fatalf(`Username "%s" expected "%t", but got "%t"`, tc.username, tc.isValid, ok)
		}
	}
}

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		email   string
		isValid bool
	}{
		// IS VALID
		{"blah@gmail.com", true},
		{"holy.moly@gmail.com", true},
		{"_somename@example.com", true},
		{"email+tag@gmail.com", true},
		{"test@d.verylongtoplevel", true},

		// NOT VALID
		{"@example.com", false},
		{"“Abc\\@def”@example.com", false},
		{"“Fred Bloggs”@example.com", false},
		{"“Joe\\Blow”@example.com", false},
		{"\\“Abc@def\\”@example.com", false},
		{"customer/department=shipping@example.com", false},
		{"$A12345@example.com", false},
		{"!def!xyz%abc@example.com", false},
		{"much.“more\\ unusual”@example.com", false},
		{"very.unusual.“@”.unusual.com@example.com", false},
		{"very.“(),:;<>[]”.VERY.“very@\\ \"very\\”.unusual@strange.example.com", false},
		{"!#$%&'*+-/=?^_`{}|~@example.com", false},
		{"Miles.O'Brian@example.com", false},
		{"postmaster@☁→❄→☃→☀→☺→☂→☹→✝.ws", false},
		{"allen@[127.0.0.1]", false},
		{"allen@[IPv6:0:0:1]", false},
		{"john@com", false},
		{"root@localhost", false},
	}

	for _, tc := range testCases {
		ok := v.IsValidEmail(tc.email)
		if ok != tc.isValid {
			t.Fatalf(`Email "%s" expected "%t", but got "%t"`, tc.email, tc.isValid, ok)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	testCases := []struct {
		password string
		isValid  bool
	}{
		{"a123456", false},  // minimum 8 characters required
		{"12346578", false}, // both number and letters required
		{"a4!@#%^$)", true}, // special characters are ok
		{"MyP4ssw0rd", true},
	}

	for _, tc := range testCases {
		ok := v.IsValidPassword(tc.password)
		if ok != tc.isValid {
			t.Fatalf(`Password "%s" expected "%t", but got "%t"`, tc.password, tc.isValid, ok)
		}
	}
}

func TestIsUniqueEmail(t *testing.T) {
	setUp()
	testCases := []struct {
		email    string
		isUnique bool
	}{
		{"system_user@localhost", false},
		{"unique_email@localhost", true},
	}

	for _, tc := range testCases {
		ok := v.IsUniqueEmail(tc.email)
		if ok != tc.isUnique {
			t.Fatalf(`Email "%s" expected "%t", but got "%t"`, tc.email, tc.isUnique, ok)
		}
	}
	tearDown()
}

func TestIsUniqueUserUsername(t *testing.T) {
	setUp()
	testCases := []struct {
		username string
		isUnique bool
	}{
		{"system_user", false},
		{"unique_user", true},
	}

	for _, tc := range testCases {
		ok := v.IsUniqueUsername(tc.username)
		if ok != tc.isUnique {
			t.Fatalf(`Username "%s" expected "%t", but got "%t"`, tc.username, tc.isUnique, ok)
		}
	}
	tearDown()
}

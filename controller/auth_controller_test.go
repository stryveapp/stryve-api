package controller_test

import (
	"net/http"
	"strings"

	"github.com/stryveapp/stryve-api/test"
)

var _ = Describe("Auth Controller", func() {

	Describe("Auth register request", func() {

		It("should return 400 Bad Request response for failing to provide all required fileds", func() {
			requestBody := `{}`
			recorder := test.JSONAPIRequest(db, "POST", "/auth/register", strings.NewReader(requestBody))

			test.JSONExpectResponseToEqual(
				recorder,
				http.StatusBadRequest,
				`{"success":false,"errors":["All feilds are required"]}`,
			)
		})

		It("should return 200 OK response for successful user registration", func() {
			requestBody := `{"username":"janedoe","email":"janedoe@example.com","password":"MyP4ssw0rd"}`
			recorder := test.JSONAPIRequest(db, "POST", "/auth/register", strings.NewReader(requestBody))

			test.JSONExpectResponseToEqual(
				recorder,
				http.StatusCreated,
				`{"success":true}`,
			)
		})

		It("should return 400 Bad Request response for providing a username of incorect length", func() {
			requestBody := `{"username":"a4","email":"janedoe2@example.com","password":"MyP4ssw0rd"}`
			recorder := test.JSONAPIRequest(db, "POST", "/auth/register", strings.NewReader(requestBody))

			test.JSONExpectResponseToEqual(
				recorder,
				http.StatusBadRequest,
				`{"success":false,"errors":["Username can only contain alpanumeric and underscore characters, and must be between 3 and 25 characters long"]}`,
			)
		})

		It("should return 400 Bad Request response for attempting to register with an existing registered username", func() {
			requestBody := `{"username":"janedoe","email":"janedoe2@example.com","password":"MyP4ssw0rd"}`
			recorder := test.JSONAPIRequest(db, "POST", "/auth/register", strings.NewReader(requestBody))

			test.JSONExpectResponseToEqual(
				recorder,
				http.StatusBadRequest,
				`{"success":false,"errors":["Your email address and/or username are already registered"]}`,
			)
		})

		It("should return 400 Bad Request response for attempting to register with an invalid email address", func() {
			requestBody := `{"username":"janedoe3","email":"janedoeATexample.com","password":"MyP4ssw0rd"}`
			recorder := test.JSONAPIRequest(db, "POST", "/auth/register", strings.NewReader(requestBody))

			test.JSONExpectResponseToEqual(
				recorder,
				http.StatusBadRequest,
				`{"success":false,"errors":["Invalid email address"]}`,
			)
		})

		It("should return 400 Bad Request response for attempting to register with an existing registered email address", func() {
			requestBody := `{"username":"janedoe2","email":"janedoe@example.com","password":"MyP4ssw0rd"}`
			recorder := test.JSONAPIRequest(db, "POST", "/auth/register", strings.NewReader(requestBody))

			test.JSONExpectResponseToEqual(
				recorder,
				http.StatusBadRequest,
				`{"success":false,"errors":["Your email address and/or username are already registered"]}`,
			)
		})

		It("should return 400 Bad Request response for attempting to register with an invlaid password", func() {
			requestBody := `{"username":"janedoe2","email":"janedoe3@example.com","password":"password"}`
			recorder := test.JSONAPIRequest(db, "POST", "/auth/register", strings.NewReader(requestBody))

			test.JSONExpectResponseToEqual(
				recorder,
				http.StatusBadRequest,
				`{"success":false,"errors":["Password must contain both numbers and letters, and be a minimum of 8 characters long"]}`,
			)
		})
	})
})

package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/mattes/migrate"
	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/controller"
	"github.com/stryveapp/stryve-api/database"
)

var (
	m  *migrate.Migrate
	db *pg.DB
	h  *controller.Handler
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Controller Test Suite")
}

var _ = BeforeSuite(func() {
	config.SetDefaultConfig()
	db = database.NewConnection("test")
	h = &controller.Handler{db}
	m = database.NewMigration("test")
	m.Up()
})

var _ = AfterSuite(func() {
	m.Down()
	db.Close()
})

func JSONAPIRequest(method string, route string, jsonData *strings.Reader) (*controller.Handler, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req, err := http.NewRequest(method, route, jsonData)
	Expect(err).ShouldNot(HaveOccurred())

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &controller.Handler{db}

	return h, c, rec
}

var _ = Describe("Auth Controller", func() {

	Describe("Auth register request", func() {

		It("should return 200 OK response for successful user registration", func() {
			requestBody := `{"username":"janedoe","email":"janedoe@example.com","password":"MyP4ssw0rd"}`
			handler, context, recorder := JSONAPIRequest("POST", "/auth/register", strings.NewReader(requestBody))

			Expect(handler.Register(context)).ShouldNot(HaveOccurred())
			Expect(recorder.Code).To(Equal(http.StatusCreated))
			Expect(recorder.Body.String()).To(Equal(`{"success":true}`))
		})

		It("should return 400 Bad Request response for failing to provide all required fileds", func() {
			requestBody := `{}`
			handler, context, recorder := JSONAPIRequest("POST", "/auth/register", strings.NewReader(requestBody))

			Expect(handler.Register(context)).ShouldNot(HaveOccurred())
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(`{"success":false,"errors":["All feilds are required"]}`))
		})

		It("should return 400 Bad Request response for providing a username of incorect length", func() {
			requestBody := `{"username":"a4","email":"janedoe2@example.com","password":"MyP4ssw0rd"}`
			handler, context, recorder := JSONAPIRequest("POST", "/auth/register", strings.NewReader(requestBody))

			Expect(handler.Register(context)).ShouldNot(HaveOccurred())
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(`{"success":false,"errors":["Username can only contain alpanumeric and underscore characters, and must be between 3 and 25 characters long"]}`))
		})

		It("should return 400 Bad Request response for attempting to register with an existing registered username", func() {
			requestBody := `{"username":"janedoe","email":"janedoe2@example.com","password":"MyP4ssw0rd"}`
			handler, context, recorder := JSONAPIRequest("POST", "/auth/register", strings.NewReader(requestBody))

			Expect(handler.Register(context)).ShouldNot(HaveOccurred())
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(`{"success":false,"errors":["Username is already taken"]}`))
		})

		It("should return 400 Bad Request response for attempting to register with an invalid email address", func() {
			requestBody := `{"username":"janedoe3","email":"janedoeATexample.com","password":"MyP4ssw0rd"}`
			handler, context, recorder := JSONAPIRequest("POST", "/auth/register", strings.NewReader(requestBody))

			Expect(handler.Register(context)).ShouldNot(HaveOccurred())
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(`{"success":false,"errors":["Invalid email address"]}`))
		})

		It("should return 400 Bad Request response for attempting to register with an existing registered email address", func() {
			requestBody := `{"username":"janedoe2","email":"janedoe@example.com","password":"MyP4ssw0rd"}`
			handler, context, recorder := JSONAPIRequest("POST", "/auth/register", strings.NewReader(requestBody))

			Expect(handler.Register(context)).ShouldNot(HaveOccurred())
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(`{"success":false,"errors":["Email address is already registered"]}`))
		})

		It("should return 400 Bad Request response for attempting to register with an invlaid password", func() {
			requestBody := `{"username":"janedoe2","email":"janedoe3@example.com","password":"password"}`
			handler, context, recorder := JSONAPIRequest("POST", "/auth/register", strings.NewReader(requestBody))

			Expect(handler.Register(context)).ShouldNot(HaveOccurred())
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(`{"success":false,"errors":["Password must contain both numbers and letters, and be a minimum of 8 characters long"]}`))
		})
	})
})

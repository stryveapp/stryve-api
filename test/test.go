package test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/gomega"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/mattes/migrate"
	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/controller"
	"github.com/stryveapp/stryve-api/database"
)

// JSONAPIRequest is a helper for running controller tests
func JSONAPIRequest(db *pg.DB, method, route string, jsonData *strings.Reader) *httptest.ResponseRecorder {
	e := echo.New()
	request, err := http.NewRequest(method, route, jsonData)
	Expect(err).NotTo(HaveOccurred())

	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	handler := &controller.Handler{DB: db}

	err = handler.Register(context)
	Expect(err).ShouldNot(HaveOccurred())

	return recorder
}

// JSONAPISetup is a helper for setting up controller tests
func JSONAPISetup() (*pg.DB, *controller.Handler, *migrate.Migrate) {
	err := config.SetDefaultConfig()
	Expect(err).NotTo(HaveOccurred())

	db := database.NewConnection("test")
	h := &controller.Handler{DB: db}

	m, err := database.NewMigration("test")
	Expect(err).NotTo(HaveOccurred())

	err = m.Up()
	Expect(err).NotTo(HaveOccurred())

	return db, h, m
}

// JSONAPITeardown is a helper for tearing down controller tests
func JSONAPITeardown(db *pg.DB, m *migrate.Migrate) {
	err := m.Down()
	Expect(err).NotTo(HaveOccurred())
	err = db.Close()
	Expect(err).NotTo(HaveOccurred())
}

// JSONExpectResponseToEqual inspects and validated the JSON response
func JSONExpectResponseToEqual(recorder *httptest.ResponseRecorder, codeExpected int, JSONExpected string) {
	Expect(recorder.Code).To(Equal(codeExpected))
	Expect(recorder.Body.String()).To(Equal(JSONExpected))
}

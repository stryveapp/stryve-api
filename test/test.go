package test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/mattes/migrate"
	"github.com/onsi/gomega"
	"github.com/stryveapp/stryve-api/config"
	"github.com/stryveapp/stryve-api/controller"
	"github.com/stryveapp/stryve-api/database"
)

var expect = gomega.Expect
var haveOccurred = gomega.HaveOccurred
var equal = gomega.Equal

// JSONAPIRequest is a helper for running controller tests
func JSONAPIRequest(db *pg.DB, method, route string, jsonData *strings.Reader) *httptest.ResponseRecorder {
	e := echo.New()
	request, err := http.NewRequest(method, route, jsonData)
	expect(err).NotTo(haveOccurred())

	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	handler := &controller.Handler{DB: db}

	err = handler.Register(context)
	expect(err).ShouldNot(haveOccurred())

	return recorder
}

// JSONAPISetup is a helper for setting up controller tests
func JSONAPISetup() (*pg.DB, *controller.Handler, *migrate.Migrate) {
	err := config.LoadDefaultConfig()
	expect(err).NotTo(haveOccurred())

	db := database.NewConnection("test")
	h := &controller.Handler{DB: db}

	m, err := database.NewMigration("test")
	expect(err).NotTo(haveOccurred())

	err = m.Up()
	expect(err).NotTo(haveOccurred())

	return db, h, m
}

// JSONAPITeardown is a helper for tearing down controller tests
func JSONAPITeardown(db *pg.DB, m *migrate.Migrate) {
	err := m.Down()
	expect(err).NotTo(haveOccurred())
	err = db.Close()
	expect(err).NotTo(haveOccurred())
}

// JSONExpectResponseToEqual inspects and validated the JSON response
func JSONExpectResponseToEqual(recorder *httptest.ResponseRecorder, codeExpected int, JSONExpected string) {
	expect(recorder.Code).To(equal(codeExpected))
	expect(recorder.Body.String()).To(equal(JSONExpected))
}

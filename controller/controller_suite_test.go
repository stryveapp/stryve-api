package controller_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/go-pg/pg"
	"github.com/mattes/migrate"
	"github.com/stryveapp/stryve-api/controller"
	"github.com/stryveapp/stryve-api/test"
)

var (
	m  *migrate.Migrate
	db *pg.DB
	h  *controller.Handler
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	db, h, m = test.JSONAPISetup()
})

var _ = AfterSuite(func() {
	test.JSONAPITeardown(db, m)
})

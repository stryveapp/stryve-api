package controller_test

import (
	"testing"

	"github.com/go-pg/pg"
	"github.com/mattes/migrate"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/stryveapp/stryve-api/controller"
	"github.com/stryveapp/stryve-api/test"
)

var (
	m  *migrate.Migrate
	db *pg.DB
	h  *controller.Handler

	RunSpecs            = ginkgo.RunSpecs
	Fail                = ginkgo.Fail
	Describe            = ginkgo.Describe
	It                  = ginkgo.It
	BeforeSuite         = ginkgo.BeforeSuite
	AfterSuite          = ginkgo.AfterSuite
	RegisterFailHandler = gomega.RegisterFailHandler
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

package validator_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validator Suite")
}

var (
	RunSpecs            = ginkgo.RunSpecs
	Fail                = ginkgo.Fail
	Describe            = ginkgo.Describe
	It                  = ginkgo.It
	BeforeSuite         = ginkgo.BeforeSuite
	AfterSuite          = ginkgo.AfterSuite
	RegisterFailHandler = gomega.RegisterFailHandler
	Expect              = gomega.Expect
	Equal               = gomega.Equal
)

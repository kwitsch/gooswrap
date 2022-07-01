package ioutil_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIoutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ioutil Suite")
}

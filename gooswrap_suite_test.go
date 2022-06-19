package gooswrap_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGooswrap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gooswrap Suite")
}

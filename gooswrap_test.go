package gooswrap_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kwitsch/gooswrap"
)

var _ = Describe("Gooswrap", func() {
	When("Wrapper in OS mode", func() {
		BeforeEach(func() {
			gooswrap.ToOs()
		})
		It("is not virtual", func() {
			v := gooswrap.Wrapper.IsVirtual()
			Expect(v).Should(BeFalse())
		})
		It("has no virtual data", func() {
			Expect(gooswrap.Wrapper.Virtual).Should(BeNil())
		})
	})
	When("Wrapper in Virtual mode", func() {
		BeforeEach(func() {
			gooswrap.NewVirtual()
		})
		It("is virtual", func() {
			v := gooswrap.Wrapper.IsVirtual()
			Expect(v).Should(BeTrue())
		})
		It("has virtual data", func() {
			Expect(gooswrap.Wrapper.Virtual).ShouldNot(BeNil())
		})
	})
})

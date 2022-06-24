package os_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kwitsch/gooswrap"

	oos "os"

	"github.com/kwitsch/gooswrap/os"
)

const (
	testKey1   string = "tk1"
	testValue1 string = "tv1"
	testKey2   string = "tk2"
	testValue2 string = "tv2"
)

var _ = Describe("Environment functions", func() {
	var (
		err error
		ok  bool
		val string
	)

	When("Wrapper in OS mode", func() {
		BeforeEach(func() {
			gooswrap.ToOs()
			err = oos.Setenv(testKey1, testValue1)
			Expect(err).Should(BeNil())
		})
		It("Clearenv", func() {
			Expect(oos.Environ()).ShouldNot(HaveLen(0))
			os.Clearenv()
			Expect(oos.Environ()).Should(HaveLen(0))
		})
		It("Environ", func() {
			Expect(os.Environ()).ShouldNot(HaveLen(0))
		})
		It("ExpandEnv", func() {
			val = os.ExpandEnv("$" + testKey1)
			Expect(val).Should(Equal(testValue1))
		})
		It("Getenv", func() {
			Expect(os.Getenv(testKey1)).Should(Equal(testValue1))
		})
		It("LookupEnv", func() {
			val, ok = os.LookupEnv(testKey1)
			Expect(ok).Should(BeTrue())
			Expect(val).Should(Equal(testValue1))
		})
		It("Setenv", func() {
			err = os.Setenv(testKey2, testValue2)
			Expect(err).Should(BeNil())
			_, ok = oos.LookupEnv(testKey2)
			Expect(ok).Should(BeTrue())
		})
		It("Unsetenv", func() {
			err = os.Unsetenv(testKey1)
			Expect(err).Should(BeNil())
			_, ok = oos.LookupEnv(testKey1)
			Expect(ok).Should(BeFalse())
		})
	})
	When("Wrapper in Virtual mode", func() {
		BeforeEach(func() {
			gooswrap.NewVirtual()
			gooswrap.Wrapper.Virtual.Env[testKey1] = testValue1
			Expect(gooswrap.Wrapper.Virtual.Env).ShouldNot(HaveLen(0))
		})
		It("Clearenv", func() {
			Expect(os.Environ()).ShouldNot(HaveLen(0))
			os.Clearenv()
			Expect(os.Environ()).Should(HaveLen(0))
		})
		It("Environ", func() {
			Expect(os.Environ()).ShouldNot(HaveLen(0))
		})
		It("ExpandEnv", func() {
			val = os.ExpandEnv("$" + testKey1)
			Expect(val).Should(Equal(testValue1))
		})
		It("LookupEnv", func() {
			val, ok = os.LookupEnv(testKey1)
			Expect(ok).Should(BeTrue())
			Expect(val).Should(Equal(testValue1))
		})
		It("Setenv", func() {
			err = os.Setenv(testKey2, testValue2)
			Expect(err).Should(BeNil())
			_, ok = gooswrap.Wrapper.Virtual.Env[testKey2]
			Expect(ok).Should(BeTrue())
		})
		It("Unsetenv", func() {
			err = os.Unsetenv(testKey1)
			Expect(err).Should(BeNil())
			_, ok = os.LookupEnv(testKey1)
			Expect(ok).Should(BeFalse())
		})
	})
})

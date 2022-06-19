package gooswrap_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	oos "os"

	"github.com/kwitsch/gooswrap"
)

const (
	testHostname string = "test"
	rootPath     string = "/"
	envKey       string = "key"
	envValue     string = "value"
)

var _ = Describe("Gooswrap", func() {
	var (
		err error
	)
	When("Wrapper in OS mode", func() {
		BeforeEach(func() {
			gooswrap.ToOs()
		})
		It("has correct working directory", func() {
			wd, err := oos.Getwd()
			Expect(err).Should(BeNil())
			Expect(gooswrap.Wrapper.WorkingDirectory).Should(Equal(wd))
		})
		It("is not virtual", func() {
			v := gooswrap.Wrapper.IsVirtual()
			Expect(v).Should(BeFalse())
		})
		It("has no virtual data", func() {
			Expect(gooswrap.Wrapper.Virtual).Should(BeNil())
		})
		It("can not execute functions for virtual", func() {
			err = gooswrap.Wrapper.Virtual.SetHostname(testHostname)
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))

			err = gooswrap.Wrapper.Virtual.SyncEnv()
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))

			err = gooswrap.Wrapper.Virtual.InitDirectories()
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))
		})
	})
	When("Wrapper in Virtual mode", func() {
		BeforeEach(func() {
			gooswrap.NewVirtual()
		})
		It("has correct working directory", func() {
			Expect(gooswrap.Wrapper.WorkingDirectory).Should(Equal(rootPath))
		})
		It("is virtual", func() {
			v := gooswrap.Wrapper.IsVirtual()
			Expect(v).Should(BeTrue())
		})
		It("has virtual data", func() {
			Expect(gooswrap.Wrapper.Virtual).ShouldNot(BeNil())
		})
		It("can execute SetHostname", func() {
			err = gooswrap.Wrapper.Virtual.SetHostname(testHostname)
			Expect(err).Should(BeNil())

			hn, err := gooswrap.Wrapper.Virtual.Hostname()
			Expect(err).Should(BeNil())
			Expect(hn).Should(Equal(testHostname))
		})
		It("can execute SyncEnv", func() {
			Expect(gooswrap.Wrapper.Virtual.Env).Should(HaveLen(0))

			oos.Clearenv()
			Expect(oos.Environ()).Should(HaveLen(0))

			err = oos.Setenv(envKey, envValue)
			Expect(err).Should(BeNil())
			Expect(oos.Environ()).Should(HaveLen(1))

			err = gooswrap.Wrapper.Virtual.SyncEnv()
			Expect(err).Should(BeNil())
			Expect(gooswrap.Wrapper.Virtual.Env).Should(HaveLen(1))
			Expect(gooswrap.Wrapper.Virtual.Env).Should(HaveKey(envKey))
			Expect(gooswrap.Wrapper.Virtual.Env[envKey]).Should(Equal(envValue))
		})
		It("can execute SyncEnv", func() {
			var exists bool

			exists, err = gooswrap.Wrapper.Util.DirExists(gooswrap.VirtualTempDir)
			Expect(err).Should(BeNil())
			Expect(exists).Should(BeFalse())

			exists, err = gooswrap.Wrapper.Util.DirExists(gooswrap.VirtualUserCacheDir)
			Expect(err).Should(BeNil())
			Expect(exists).Should(BeFalse())

			exists, err = gooswrap.Wrapper.Util.DirExists(gooswrap.VirtualUserConfigDir)
			Expect(err).Should(BeNil())
			Expect(exists).Should(BeFalse())

			exists, err = gooswrap.Wrapper.Util.DirExists(gooswrap.VirtualUserHomeDir)
			Expect(err).Should(BeNil())
			Expect(exists).Should(BeFalse())

			err = gooswrap.Wrapper.Virtual.InitDirectories()
			Expect(err).Should(BeNil())

			exists, err = gooswrap.Wrapper.Util.DirExists(gooswrap.VirtualTempDir)
			Expect(err).Should(BeNil())
			Expect(exists).Should(BeTrue())

			exists, err = gooswrap.Wrapper.Util.DirExists(gooswrap.VirtualUserCacheDir)
			Expect(err).Should(BeNil())
			Expect(exists).Should(BeTrue())

			exists, err = gooswrap.Wrapper.Util.DirExists(gooswrap.VirtualUserConfigDir)
			Expect(err).Should(BeNil())
			Expect(exists).Should(BeTrue())

			exists, err = gooswrap.Wrapper.Util.DirExists(gooswrap.VirtualUserHomeDir)
			Expect(err).Should(BeNil())
			Expect(exists).Should(BeTrue())
		})
	})
})

package gooswrap_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	oos "os"

	"github.com/kwitsch/gooswrap"
)

const (
	testId       int    = 99
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
			err = gooswrap.Wrapper.Virtual.SetEgid(testId)
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))

			err = gooswrap.Wrapper.Virtual.SetEuid(testId)
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))

			err = gooswrap.Wrapper.Virtual.SetGid(testId)
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))

			err = gooswrap.Wrapper.Virtual.SetGroups([]int{testId})
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))

			err = gooswrap.Wrapper.Virtual.SetPagesize(testId)
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))

			err = gooswrap.Wrapper.Virtual.SetPid(testId)
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))

			err = gooswrap.Wrapper.Virtual.SetPpid(testId)
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))

			err = gooswrap.Wrapper.Virtual.SetUid(testId)
			Expect(err).ShouldNot(BeNil())
			Expect(err).Should(Equal(gooswrap.ErrNotVirtual))

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
		var (
			id int
		)
		BeforeEach(func() {
			gooswrap.NewVirtual()
			id = 0
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
		It("can execute SetEgid", func() {
			id, err = gooswrap.Wrapper.Virtual.Egid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(oos.Getegid()))

			err = gooswrap.Wrapper.Virtual.SetEgid(testId)
			Expect(err).Should(BeNil())

			id, err = gooswrap.Wrapper.Virtual.Egid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(testId))
		})
		It("can execute SetEuid", func() {
			id, err = gooswrap.Wrapper.Virtual.Euid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(oos.Geteuid()))

			err = gooswrap.Wrapper.Virtual.SetEuid(testId)
			Expect(err).Should(BeNil())

			id, err = gooswrap.Wrapper.Virtual.Euid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(testId))
		})
		It("can execute SetGid", func() {
			id, err = gooswrap.Wrapper.Virtual.Gid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(oos.Getgid()))

			err = gooswrap.Wrapper.Virtual.SetGid(testId)
			Expect(err).Should(BeNil())

			id, err = gooswrap.Wrapper.Virtual.Gid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(testId))
		})
		It("can execute SetGroups", func() {
			ids, err := gooswrap.Wrapper.Virtual.Groups()
			Expect(err).Should(BeNil())
			Expect(ids).Should(HaveLen(0))

			err = gooswrap.Wrapper.Virtual.SetGroups([]int{testId})
			Expect(err).Should(BeNil())

			ids, err = gooswrap.Wrapper.Virtual.Groups()
			Expect(err).Should(BeNil())
			Expect(ids).Should(HaveLen(1))
			Expect(ids[0]).Should(Equal(testId))
		})
		It("can execute SetPagesize", func() {
			id, err = gooswrap.Wrapper.Virtual.Pagesize()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(oos.Getpagesize()))

			err = gooswrap.Wrapper.Virtual.SetPagesize(testId)
			Expect(err).Should(BeNil())

			id, err = gooswrap.Wrapper.Virtual.Pagesize()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(testId))
		})
		It("can execute SetPid", func() {
			id, err = gooswrap.Wrapper.Virtual.Pid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(oos.Getpid()))

			err = gooswrap.Wrapper.Virtual.SetPid(testId)
			Expect(err).Should(BeNil())

			id, err = gooswrap.Wrapper.Virtual.Pid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(testId))
		})
		It("can execute SetPpid", func() {
			id, err = gooswrap.Wrapper.Virtual.Ppid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(oos.Getppid()))

			err = gooswrap.Wrapper.Virtual.SetPpid(testId)
			Expect(err).Should(BeNil())

			id, err = gooswrap.Wrapper.Virtual.Ppid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(testId))
		})
		It("can execute SetUid", func() {
			id, err = gooswrap.Wrapper.Virtual.Uid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(oos.Getuid()))

			err = gooswrap.Wrapper.Virtual.SetUid(testId)
			Expect(err).Should(BeNil())

			id, err = gooswrap.Wrapper.Virtual.Uid()
			Expect(err).Should(BeNil())
			Expect(id).Should(Equal(testId))
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
	})
})

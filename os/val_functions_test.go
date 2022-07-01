package os_test

import (
	"github.com/avfs/avfs/idm/memidm"
	"github.com/avfs/avfs/vfs/memfs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kwitsch/gooswrap"

	oos "os"

	"github.com/kwitsch/gooswrap/os"
)

var _ = Describe("Value functions", func() {
	When("Wrapper in OS mode", func() {
		BeforeEach(func() {
			gooswrap.ToOs()
		})
		It("Getegid", func() {
			Expect(os.Getegid()).Should(Equal(oos.Getegid()))
		})
		It("Geteuid", func() {
			Expect(os.Geteuid()).Should(Equal(oos.Geteuid()))
		})
		It("Getgid", func() {
			Expect(os.Getgid()).Should(Equal(oos.Getgid()))
		})
		It("Getgroups", func() {
			oov, ooerr := oos.Getgroups()
			ov, oerr := os.Getgroups()
			Expect(oerr).Should(Equal(ooerr))
			Expect(ov).Should(Equal(oov))
		})
		It("Getpagesize", func() {
			Expect(os.Getpagesize()).Should(Equal(oos.Getpagesize()))
		})
		It("Getpid", func() {
			Expect(os.Getpid()).Should(Equal(oos.Getpid()))
		})
		It("Getppid", func() {
			Expect(os.Getppid()).Should(Equal(oos.Getppid()))
		})
		It("Getuid", func() {
			Expect(os.Getuid()).Should(Equal(oos.Getuid()))
		})
		It("Getwd", func() {
			oov, ooerr := oos.Getwd()
			ov, oerr := os.Getwd()
			Expect(ooerr).Should(BeNil())
			Expect(oerr).Should(BeNil())
			Expect(ov).Should(Equal(oov))
		})
		It("Hostname", func() {
			oov, ooerr := oos.Hostname()
			ov, oerr := os.Hostname()
			Expect(ooerr).Should(BeNil())
			Expect(oerr).Should(BeNil())
			Expect(ov).Should(Equal(oov))
		})
	})

	When("Wrapper in Virtual mode", func() {
		const (
			egid     int    = 991
			euid     int    = 992
			gid      int    = 993
			group    int    = 994
			pagesize int    = 995
			pid      int    = 996
			ppid     int    = 997
			uid      int    = 998
			hostname string = "virtual"
		)
		var (
			err error
		)
		BeforeEach(func() {
			gooswrap.NewVirtual()
			err = gooswrap.Wrapper.Virtual.SetEgid(egid)
			Expect(err).Should(BeNil())
			err = gooswrap.Wrapper.Virtual.SetEuid(euid)
			Expect(err).Should(BeNil())
			err = gooswrap.Wrapper.Virtual.SetGid(gid)
			Expect(err).Should(BeNil())
			err = gooswrap.Wrapper.Virtual.SetGroups([]int{group})
			Expect(err).Should(BeNil())
			err = gooswrap.Wrapper.Virtual.SetPagesize(pagesize)
			Expect(err).Should(BeNil())
			err = gooswrap.Wrapper.Virtual.SetPid(pid)
			Expect(err).Should(BeNil())
			err = gooswrap.Wrapper.Virtual.SetPpid(ppid)
			Expect(err).Should(BeNil())
			err = gooswrap.Wrapper.Virtual.SetUid(uid)
			Expect(err).Should(BeNil())
		})
		It("Getegid", func() {
			Expect(os.Getegid()).Should(Equal(egid))
		})
		It("Geteuid", func() {
			Expect(os.Geteuid()).Should(Equal(euid))
		})
		It("Getgid", func() {
			Expect(os.Getgid()).Should(Equal(gid))
		})
		It("Getgroups", func() {
			ov, oerr := os.Getgroups()
			Expect(oerr).Should(BeNil())
			Expect(ov).Should(HaveLen(1))
			Expect(ov[0]).Should(Equal(group))
		})
		It("Getpagesize", func() {
			Expect(os.Getpagesize()).Should(Equal(pagesize))
		})
		It("Getpid", func() {
			Expect(os.Getpid()).Should(Equal(pid))
		})
		It("Getppid", func() {
			Expect(os.Getppid()).Should(Equal(ppid))
		})
		It("Getuid", func() {
			Expect(os.Getuid()).Should(Equal(uid))
		})
		It("Getwd", func() {
			av, aerr := memfs.New(memfs.WithMainDirs(), memfs.WithIdm(memidm.New())).Getwd()
			Expect(aerr).Should(BeNil())
			ov, oerr := os.Getwd()
			Expect(oerr).Should(BeNil())
			Expect(ov).Should(Equal(av))
		})
		It("Hostname", func() {
			ov, oerr := os.Hostname()
			Expect(oerr).Should(BeNil())
			Expect(ov).Should(Equal(hostname))
		})
	})
})

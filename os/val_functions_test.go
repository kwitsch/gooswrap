package os_test

import (
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
			wd       string = "/"
			hostname string = "virtual"
		)
		BeforeEach(func() {
			gooswrap.NewVirtual()
			gooswrap.Wrapper.Virtual.SetEgid(egid)
			gooswrap.Wrapper.Virtual.SetEuid(euid)
			gooswrap.Wrapper.Virtual.SetGid(gid)
			gooswrap.Wrapper.Virtual.SetGroups([]int{group})
			gooswrap.Wrapper.Virtual.SetPagesize(pagesize)
			gooswrap.Wrapper.Virtual.SetPid(pid)
			gooswrap.Wrapper.Virtual.SetPpid(ppid)
			gooswrap.Wrapper.Virtual.SetUid(uid)
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
			ov, oerr := os.Getwd()
			Expect(oerr).Should(BeNil())
			Expect(ov).Should(Equal(wd))
		})
		It("Hostname", func() {
			ov, oerr := os.Hostname()
			Expect(oerr).Should(BeNil())
			Expect(ov).Should(Equal(hostname))
		})
	})
})

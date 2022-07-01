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

var _ = Describe("Directory functions", func() {
	When("Wrapper in OS mode", func() {
		var (
			oov string
			ov  string
		)
		BeforeEach(func() {
			gooswrap.ToOs()
		})
		It("TempDir", func() {
			Expect(os.TempDir()).Should(Equal(oos.TempDir()))
		})
		It("UserCacheDir", func() {
			oov, _ = oos.UserCacheDir()
			ov, _ = os.UserCacheDir()
			Expect(ov).Should(Equal(oov))
		})
		It("UserConfigDir", func() {
			oov, _ = oos.UserConfigDir()
			ov, _ = os.UserConfigDir()
			Expect(ov).Should(Equal(oov))
		})
		It("UserHomeDir", func() {
			oov, _ = oos.UserHomeDir()
			ov, _ = os.UserHomeDir()
			Expect(ov).Should(Equal(oov))
		})
	})
	When("Wrapper in Virtual mode", func() {
		var (
			ov  string
			err error
		)
		BeforeEach(func() {
			gooswrap.NewVirtual()
		})
		It("TempDir", func() {
			tmpFs := memfs.New(memfs.WithMainDirs(), memfs.WithIdm(memidm.New()))
			Expect(os.TempDir()).Should(Equal(tmpFs.TempDir()))
		})
		It("UserCacheDir", func() {
			ov, err = os.UserCacheDir()
			Expect(err).Should(BeNil())
			Expect(ov).Should(Equal(gooswrap.VirtualUserCacheDir))
		})
		It("UserConfigDir", func() {
			ov, err = os.UserConfigDir()
			Expect(err).Should(BeNil())
			Expect(ov).Should(Equal(gooswrap.VirtualUserConfigDir))
		})
		It("UserHomeDir", func() {
			ov, err = os.UserHomeDir()
			Expect(err).Should(BeNil())
			Expect(ov).Should(Equal(gooswrap.VirtualUserHomeDir))
		})
	})
})

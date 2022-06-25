# Go Os Wrapper

gooswrap is a thin wrapper around the os and io/ioutil packages.

The wrapper provides a virtualisation layer for int testing in an automated environment.

## Production usage

Just replace the os and io/ioutil packages to the ones provided in this package.

```golang
import (
    "os"
    "io/ioutil"
)
```

To

```golang
import (
    "github.com/kwitsch/gooswrap/os"
    "github.com/kwitsch/gooswrap/ioutil"
)
```

## Test usage

A virtual environment can be initialized with `gooswrap.NewVirtual()`.  
After that several values specified can be set in the virtual environment.  
Some methodes will return `ErrNotSupported` in virtual mode as not all functions can be virtulaized.  

Ginkgo example:

```golang
When("Set values in virtual environment", func() {
    const (
        egid     int    = 991
        euid     int    = 992
        gid      int    = 993
        group    int    = 994
        pagesize int    = 995
        pid      int    = 996
        ppid     int    = 997
        uid      int    = 998
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
```  

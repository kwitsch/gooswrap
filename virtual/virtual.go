package virtual

import (
	oos "os"
	"strings"

	. "github.com/kwitsch/gooswrap"
)

func SyncEnv() error {
	return onlyWhenVirtual(func() {
		oenv := oos.Environ()
		for _, ec := range oenv {
			es := strings.Split(ec, "=")
			if len(es) == 2 {
				Wrapper.Data.Env[es[0]] = es[1]
			}
		}
	})
}

func SetHostname(hostname string) error {
	return onlyWhenVirtual(func() {
		Wrapper.Data.Hostname = hostname
	})
}

func onlyWhenVirtual(action func()) error {
	if IsVirtual() {
		action()
		return nil
	}

	return ErrNotVirtual
}

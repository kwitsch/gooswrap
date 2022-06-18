package virtual

import (
	oos "os"
	"strings"

	. "github.com/kwitsch/gooswrap"
)

func InitEnv() {
	if IsVirtual() {
		oenv := oos.Environ()
		for _, ec := range oenv {
			es := strings.Split(ec, "=")
			if len(es) == 2 {
				Wrapper.Data.Env[es[0]] = es[1]
			}
		}
	}
}

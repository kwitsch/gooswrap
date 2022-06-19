package os

import (
	"fmt"
	oos "os"

	. "github.com/kwitsch/gooswrap"
)

// Clearenv deletes all environment variables.
func Clearenv() {
	if !Wrapper.IsVirtual() {
		oos.Clearenv()
	} else {
		Wrapper.Virtual.Env = make(map[string]string)
	}
}

// Environ returns a copy of strings representing the environment,
// in the form "key=value".
func Environ() []string {
	var res []string
	if !Wrapper.IsVirtual() {
		res = oos.Environ()
	} else {
		res = make([]string, len(Wrapper.Virtual.Env))
		i := 0
		for k, v := range Wrapper.Virtual.Env {
			res[i] = fmt.Sprintf("%s=%s", k, v)
			i++
		}
	}
	return res
}

// ExpandEnv replaces ${var} or $var in the string according to the values
// of the current environment variables. References to undefined
// variables are replaced by the empty string.
func ExpandEnv(s string) string {
	return Expand(s, Getenv)
}

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
// To distinguish between an empty value and an unset value, use LookupEnv.
func Getenv(key string) string {
	if !Wrapper.IsVirtual() {
		return oos.Getenv(key)
	} else {
		if v, ok := Wrapper.Virtual.Env[key]; ok {
			return v
		} else {
			return ""
		}
	}
}

// LookupEnv retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will
// be false.
func LookupEnv(key string) (string, bool) {
	if !Wrapper.IsVirtual() {
		return oos.LookupEnv(key)
	} else {
		v, ok := Wrapper.Virtual.Env[key]
		return v, ok
	}
}

// Setenv sets the value of the environment variable named by the key.
// It returns an error, if any.
func Setenv(key, value string) error {
	if !Wrapper.IsVirtual() {
		return oos.Setenv(key, value)
	} else {
		Wrapper.Virtual.Env[key] = value
		return nil
	}
}

// Unsetenv unsets a single environment variable.
func Unsetenv(key string) error {
	if !Wrapper.IsVirtual() {
		return oos.Unsetenv(key)
	} else {
		delete(Wrapper.Virtual.Env, key)
		return nil
	}
}

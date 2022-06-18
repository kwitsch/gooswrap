package os

import (
	oos "os"
)

// TempDir returns the default directory to use for temporary files.
//
// On Unix systems, it returns $TMPDIR if non-empty, else /tmp.
// On Windows, it uses GetTempPath, returning the first non-empty
// value from %TMP%, %TEMP%, %USERPROFILE%, or the Windows directory.
// On Plan 9, it returns /tmp.
//
// The directory is neither guaranteed to exist nor have accessible
// permissions.
func TempDir() string {
	return oos.TempDir()
}

// UserCacheDir returns the default root directory to use for user-specific
// cached data. Users should create their own application-specific subdirectory
// within this one and use that.
//
// On Unix systems, it returns $XDG_CACHE_HOME as specified by
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
// non-empty, else $HOME/.cache.
// On Darwin, it returns $HOME/Library/Caches.
// On Windows, it returns %LocalAppData%.
// On Plan 9, it returns $home/lib/cache.
//
// If the location cannot be determined (for example, $HOME is not defined),
// then it will return an error.
func UserCacheDir() (string, error) {
	return oos.UserCacheDir()
}

// UserConfigDir returns the default root directory to use for user-specific
// configuration data. Users should create their own application-specific
// subdirectory within this one and use that.
//
// On Unix systems, it returns $XDG_CONFIG_HOME as specified by
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
// non-empty, else $HOME/.config.
// On Darwin, it returns $HOME/Library/Application Support.
// On Windows, it returns %AppData%.
// On Plan 9, it returns $home/lib.
//
// If the location cannot be determined (for example, $HOME is not defined),
// then it will return an error.
func UserConfigDir() (string, error) {
	return oos.UserConfigDir()
}

// UserHomeDir returns the current user's home directory.
//
// On Unix, including macOS, it returns the $HOME environment variable.
// On Windows, it returns %USERPROFILE%.
// On Plan 9, it returns the $home environment variable.
func UserHomeDir() (string, error) {
	return oos.UserHomeDir()
}

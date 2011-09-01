package errforce

import (
	"os"
)

// Return true if the given error is ENOENT and false otherwise.
func IsENOENT(err os.Error) bool {
	if perr, ok := err.(*os.PathError); ok {
		err = perr.Error
	}
	if errno, ok := err.(os.Errno); ok {
		return os.ENOENT == errno
	}
	return false
}

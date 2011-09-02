// Utilities for working effectively with Go errors and UNIX errnos.
package errforce

import (
	"compress/flate"
	"csv"
	"exec"
	"image/png"
	"json"
	"net"
	"os"
	"strconv"
	"url"
	"websocket"
)

// Return true if err represents the same errno as refErr, which should be
// one of the os.E* constants such as os.EPERM or os.ENOENT.
func Is(refErr, err os.Error) bool {

	// It'd be nice to use a big type switch here to but the compiler
	// complains about os.Error not having a member named Error.  This seems
	// silly, since the whole point of a type switch is to operate with the
	// correct runtime type rather than the interface type.
	if werr, ok := err.(*csv.ParseError); ok { err = werr.Error }
	if werr, ok := err.(*exec.Error); ok { err = werr.Error }
	if werr, ok := err.(*flate.ReadError); ok { err = werr.Error }
	if werr, ok := err.(*flate.WriteError); ok { err = werr.Error }
	if werr, ok := err.(*json.MarshalerError); ok { err = werr.Error }
	if werr, ok := err.(*net.DNSConfigError); ok { err = werr.Error }
	if werr, ok := err.(*net.OpError); ok { err = werr.Error }
	if werr, ok := err.(*os.LinkError); ok { err = werr.Error }
	if werr, ok := err.(*os.PathError); ok { err = werr.Error }
	if werr, ok := err.(png.IDATDecodingError); ok { err = werr.Err }
	if werr, ok := err.(*strconv.NumError); ok { err = werr.Error }
	if werr, ok := err.(*url.Error); ok { err = werr.Error }
	if werr, ok := err.(*websocket.DialError); ok { err = werr.Error }

	// In practice, it doesn't matter that os.SyscallError doesn't appear
	// here because system calls themselves return integral errno, which can
	// be rolled up into an os.Errno, not an os.SyscallError.

	if errno, ok := err.(os.Errno); ok { return refErr == errno }
	return false
}

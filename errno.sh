cat <<EOF
package errforce

import "os"

EOF

awk '
/ Error = Errno\(syscall\.E/ {
	print "func Is" $1 "(err os.Error) bool { return Is(os." $1 ", err) }"
}
' <"$GOROOT/src/pkg/os/error_posix.go"

//go:build unix && !(linux || darwin || freebsd || netbsd || dragonfly)

package platform

import "syscall"

func MprotectRX(b []byte) error {
	return syscall.ENOTSUP
}

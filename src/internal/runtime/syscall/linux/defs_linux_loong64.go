// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linux

const (
	SYS_CLOSE         = 57
	SYS_EPOLL_CREATE1 = 20
	SYS_EPOLL_CTL     = 21
	SYS_EPOLL_PWAIT   = 22
	SYS_FCNTL         = 25
	SYS_PRCTL         = 167
	SYS_MPROTECT      = 226
	SYS_EPOLL_PWAIT2  = 441
	SYS_EVENTFD2      = 19
	SYS_OPENAT        = 56
	SYS_PREAD64       = 67
	SYS_READ          = 63

	EFD_NONBLOCK = 0x800

	O_LARGEFILE = 0x0
)

type EpollEvent struct {
	Events    uint32
	pad_cgo_0 [4]byte
	Data      [8]byte // unaligned uintptr
}

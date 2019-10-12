// Copyright 2016 Tim O'Brien. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package javabind

import "golang.org/x/sys/unix"

func GetThreadId() int {
	r0, _, _ := unix.RawSyscall(unix.SYS_THREAD_SELFID, 0, 0, 0)
	return int(r0)
}

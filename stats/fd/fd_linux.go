/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

// +build linux

package fd

import (
	"github.com/anatma/procfs"
)

type FileDescriptor struct {
	PID int
	FDs map[string]string
}

func (m *FD) FDs() (fds []FileDescriptor) {
	fs, err := procfs.NewFS(procfs.DefaultMountPoint)
	if err != nil {
		return nil
	}

	procs, err := fs.AllProcs()
	if err != nil {
		return nil
	}

	for _, i := range procs {
		fd, _ := i.NewFD()

		if len(fd) != 0 {
			fds = append(fds, FileDescriptor{
				PID: i.PID,
				FDs: fd,
			})
		}
	}

	return fds
}

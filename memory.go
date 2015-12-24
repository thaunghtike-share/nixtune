/*
 * Anatma Autotune - Kernel Autotuning
 *
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"github.com/abhiyerra/procfs"
)

type Memory struct {
	Ram  int64
	Swap int64
}

func ComputeMemory() (m *Memory) {
	fs, err := procfs.NewFS(procfs.DefaultMountPoint)
	if err != nil {
		return
	}

	meminfo, err := fs.NewMeminfo()
	if err != nil {
		return
	}

	m.Ram = meminfo.MemTotal - meminfo.MemFree
	m.Ram = meminfo.SwapTotal - meminfo.SwapFree

	return
}

func MemoryWorker() {
	for {

	}
}

/*

https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html

Look to see how memory changes over time.

 - How many malloc allocations are done.
 - How quickly does the

*/

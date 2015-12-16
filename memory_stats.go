/*
 * Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package main

import (
	"github.com/abhiyerra/procfs"
)

type MemoryStats struct {
	Total int64
	Free  int64
}

func (n *MemoryStats) ParseArgs(args []string) {

}

func (m *MemoryStats) Run() error {
	fs, err := procfs.NewFS(procfs.DefaultMountPoint)
	if err != nil {
		return err
	}

	meminfo, err := fs.NewMeminfo()
	if err != nil {
		return err
	}

	m.Total = meminfo.MemTotal
	m.Free = meminfo.MemFree

	return nil
}

func NewMemoryStats() *MemoryStats {
	return &MemoryStats{}
}

/*
https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html

Look to see how memory changes over time.

 - How many malloc allocations are done.
 - How quickly does the

*/

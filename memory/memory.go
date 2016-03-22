/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package memory

import (
	"github.com/acksin/procfs"
)

/*
https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html

Look to see how memory changes over time.

 - How many malloc allocations are done.
 - How quickly does the

 Documentation for Linux Memory
 - http://superuser.com/questions/521551/cat-proc-meminfo-what-do-all-those-numbers-mean

*/

// Memory returns the memory usage of the system.
type Memory struct {
	// Unit is the metric used for all the numbers returned. E.g. kb, mb, gb.
	Unit string

	// Total is the total memory of the system.
	Total int64
	// Free is the total free memory in the system.
	Free int64
	// Cached is the total memory that is cached.
	Cached int64

	// Swap is information on swap usage on the system.
	Swap struct {
		// Total swap that is available.
		Total int64
		// Free swap space.
		Free int64
		// Used swap space.
		Used int64
	}
}

func (m *Memory) mem(meminfo *procfs.Meminfo) {
	m.Total = meminfo.MemTotal
	m.Free = meminfo.MemFree
	m.Cached = meminfo.Cached
}

func (m *Memory) swap(meminfo *procfs.Meminfo) {
	m.Swap.Used = meminfo.SwapCached
	m.Swap.Total = meminfo.SwapTotal
	m.Swap.Free = meminfo.SwapFree
}

// New returns a Memory object representing system memory information.
func New() *Memory {
	m := &Memory{}

	meminfo, err := procfs.NewMeminfo()
	if err != nil {
		return nil
	}

	m.Unit = "kb"

	m.swap(&meminfo)
	m.mem(&meminfo)

	return m
}

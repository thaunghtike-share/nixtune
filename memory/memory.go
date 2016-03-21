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
type Memory struct {
	Unit string

	Total  int64
	Free   int64
	Cached int64

	Swap struct {
		Total int64
		Free  int64
		Used  int64
	}
}

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

/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package memory

import (
	"strconv"

	"github.com/anatma/procfs"
)

type ProcessMemory struct {
	proc procfs.Proc

	Swap struct {
		Size     int64  `json:",omitempty"`
		SizeUnit string `json:",omitempty"`
	} `json:",omitempty"`
}

func NewProcess(proc procfs.Proc) *ProcessMemory {
	pm := &ProcessMemory{proc: proc}
	pm.getSwap()

	return pm
}

func (m *ProcessMemory) getSwap() {
	ps, _ := m.proc.NewStatus()
	if ps.VmSwap == "" {
		return
	}

	v, err := strconv.ParseInt(ps.VmSwap, 10, 64)
	if err != nil {
		return
	}

	// It's swapping.
	if v > 0 {
		m.Swap.Size = v
		m.Swap.SizeUnit = "kb"
	}
}

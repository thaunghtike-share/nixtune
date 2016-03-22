/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package memory

import (
	"strconv"

	"github.com/acksin/procfs"
)

// ProcessMemory returns memory usage of a Process
type ProcessMemory struct {
	// Swap usage of the Process
	Swap struct {
		// Size of the swap used
		Size int64 `json:",omitempty"`
		// Unit metric used
		Unit string `json:",omitempty"`
	} `json:",omitempty"`

	proc procfs.Proc
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
		m.Swap.Unit = "kb"
	}
}

// NewProcess returns memory information of a Linux Process
func NewProcess(proc procfs.Proc) *ProcessMemory {
	pm := &ProcessMemory{proc: proc}
	pm.getSwap()

	return pm
}

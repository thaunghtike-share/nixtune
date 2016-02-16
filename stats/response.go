/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package stats

import (
	"github.com/acksin/autotune/stats/fd"
	"github.com/acksin/autotune/stats/memory"
	"github.com/acksin/autotune/stats/network"
)

type Process struct {
	Exe    string
	PID    int
	Memory *memory.ProcessMemory `json:",omitempty"`
	FD     fd.ProcessFD          `json:",omitempty"`
}

type Response struct {
	System struct {
		Memory  *memory.Memory
		Network *network.Network
	}

	Processes []Process
}

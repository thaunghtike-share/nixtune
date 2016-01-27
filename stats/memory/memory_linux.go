/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

// +build linux

package memory

import (
	"log"

	"github.com/anatma/procfs"
)

// Documentation for Linux Memory
// - http://superuser.com/questions/521551/cat-proc-meminfo-what-do-all-those-numbers-mean

func (m *Memory) Swapping() bool {
	meminfo, err := procfs.NewMeminfo()
	if err != nil {
		log.Printf("could not get meminfo: %s", err)
		return false
	}

	if meminfo.SwapCached > 0 {
		return true
	}

	return false
}

func (m *Memory) SwappingProcesses() map[string]bool {
	swappingProcesses := make(map[string]bool)

	return swappingProcesses
}

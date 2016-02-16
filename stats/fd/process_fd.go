/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package fd

import (
	"github.com/acksin/procfs"
)

type ProcessFD map[string]string

func NewProcess(proc procfs.Proc) ProcessFD {
	fd, err := proc.NewFD()
	if err != nil {

	}

	return ProcessFD(fd)
}

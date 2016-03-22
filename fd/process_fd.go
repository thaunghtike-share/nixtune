/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package fd

import (
	"github.com/acksin/procfs"
)

// ProcessFD is a map of the file descriptor number to the descriptor
// object that it is pointing to.
type ProcessFD map[string]string

// NewProcess returns the file descriptors for a single Linux process.
func NewProcess(proc procfs.Proc) ProcessFD {
	fd, err := proc.NewFD()
	if err != nil {
		return nil
	}

	return ProcessFD(fd)
}

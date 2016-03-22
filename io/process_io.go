/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package io

import (
	"github.com/acksin/procfs"
)

type ProcessIO struct {
	FD ProcessFD
}

func (p *ProcessIO) fileDescriptors(proc *procfs.Proc) {
	fd, err := proc.NewFD()
	if err != nil {
		return
	}
	p.FD = ProcessFD(fd)
}

// NewProcess returns the file descriptors for a single Linux process.
func NewProcess(proc procfs.Proc) *ProcessIO {
	p := &ProcessIO{}
	p.fileDescriptors(&proc)

	return p
}

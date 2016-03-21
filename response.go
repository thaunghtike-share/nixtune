/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"github.com/acksin/procfs"
	"github.com/acksin/strum/fd"
	"github.com/acksin/strum/memory"
	"github.com/acksin/strum/network"
)

// OutputType is the formatted output of the command.
type OutputType string

// Currently available output types.
const (
	JsonType  OutputType = "json"
	Flattened            = "flattened"
	Human                = "human"
)

type Stats struct {
	Output string
}

func (n *Stats) processes() procfs.Procs {
	fs, err := procfs.NewFS(procfs.DefaultMountPoint)
	if err != nil {
		return nil
	}

	procs, err := fs.AllProcs()
	if err != nil {
		return nil
	}

	return procs
}

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

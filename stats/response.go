/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package stats

import (
	"encoding/json"

	"github.com/acksin/procfs"
	"github.com/acksin/strum/fd"
	"github.com/acksin/strum/memory"
	"github.com/acksin/strum/network"
)

type Process struct {
	Exe    string
	PID    int
	Memory *memory.ProcessMemory `json:",omitempty"`
	FD     fd.ProcessFD          `json:",omitempty"`
}

type Stats struct {
	System struct {
		Memory  *memory.Memory
		Network *network.Network
	}

	Processes []Process
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

func (n *Stats) Json() string {
	js, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return ""
	}

	return string(js)
}

func New() (s *Stats) {
	s = &Stats{}

	s.System.Memory = memory.New()

	for _, proc := range s.processes() {
		exe, err := proc.Executable()
		if err != nil || exe == "" {
			status, _ := proc.NewStatus()
			exe = status.Name
		}

		p := Process{
			Exe:    exe,
			PID:    proc.PID,
			Memory: memory.NewProcess(proc),
			FD:     fd.NewProcess(proc),
		}

		s.Processes = append(s.Processes, p)
	}

	return s
}

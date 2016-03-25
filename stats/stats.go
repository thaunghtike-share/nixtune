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
	"fmt"
	"sort"

	"github.com/abhiyerra/gojsonexplode"

	"github.com/acksin/procfs"
	"github.com/acksin/strum/io"
	"github.com/acksin/strum/memory"
	"github.com/acksin/strum/network"
)

// Process is information about a Linux process
type Process struct {
	// Exe is the executable that is running.
	Exe string
	// PID of the process
	PID int
	// Memory stats of the process
	Memory *memory.ProcessMemory
	// IO contains information about the IO of the machine.
	IO *io.ProcessIO
}

// Stats contains both the system and process statistics.
type Stats struct {
	// System specific information
	System struct {
		// Memory stats of the system
		Memory *memory.Memory
		// Network stats of the system
		Network *network.Network
	}

	// Processes are the process information of the system
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

// JSON returns JSON string of Stats
func (n *Stats) JSON() string {
	js, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return ""
	}

	return string(js)
}

// Flat returns a flattened results.
func (n *Stats) Flat() string {
	o, err := gojsonexplode.Explodejsonstr(n.JSON(), ".")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var out map[string]interface{}

	err = json.Unmarshal([]byte(o), &out)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var o2 string
	var keys []string

	for k := range out {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		o2 += fmt.Sprintf("%s = %v\n", k, out[k])
	}

	return o2
}

func (n *Stats) containsPid(pids []int, proc procfs.Proc) bool {
	for _, pid := range pids {
		if proc.PID == pid {
			return true
		}
	}

	return false
}

// New returns stats of the machine with pids filtering for
// processes. If pids are empty then it returns all process stats.
func New(pids []int) (s *Stats) {
	s = &Stats{}

	s.System.Memory = memory.New()
	s.System.Network = network.New()

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
			IO:     io.NewProcess(proc),
		}

		if len(pids) == 0 || s.containsPid(pids, proc) {
			s.Processes = append(s.Processes, p)
		}
	}

	return s
}

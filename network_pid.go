/*
 * Anatma Autotune - Kernel Autotuning
 *
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"flag"
	"fmt"
	"github.com/abhiyerra/procfs"
	"os"
)

type NetworkPid struct {
	Duration int
}

// Return the PID that has the most amount of network calls on the
// server. We want to guess it based on TCP/UDP utilization.
func (n *NetworkPid) MainPid() string {
	uidMode := make(map[string]uint)
	netTcp, _ := procfs.NewNetTCP()

	for _, i := range netTcp {
		uidMode[i.UID]++
	}

	var numUidVal uint = 0
	mainUid := ""

	for k, v := range uidMode {
		if v > numUidVal {
			mainUid = k
		}
	}

	return mainUid
}

func (n *NetworkPid) ParseArgs(args []string) {
	flags := flag.NewFlagSet(CmdName, flag.ContinueOnError)

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}
}

func (n *NetworkPid) Run() error {
	fmt.Println(n.MainPid())

	return nil
}

func NewNetworkPid() *NetworkPid {
	return &NetworkPid{}
}

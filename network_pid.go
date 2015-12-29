/*
 * Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
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
	flags := flag.NewFlagSet(subCmd("network", "pid"), flag.ContinueOnError)

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

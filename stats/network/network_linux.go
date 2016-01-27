/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package stats

import (
	"fmt"
	"time"

	"flag"
	"github.com/abhiyerra/procfs"
	"github.com/montanaflynn/stats"
	"os"
)

/*

TODO:
Network Access ALgorithm for Anatma Autotune

To be able to do a profile on network throughput we need to first take
a metric of teh connection.

 - Need to look at network connections over a period of time.
 - See if the connections are to the same pid of different ones.
 - See how many timeouts there are.
 - Look at how many connections there are that are open and how tey are operating.
 - See if there is an increase see the throughput as it grows over time.
 - Need to profile the machine over time. Over a long period of time.
*/

// getNetworkSettings()

type Network struct {
	CmdName string

	Duration     int
	NetSockstats []procfs.NetSockstat

	Profile struct {
		// Only connects to a limited number of remote
		// addresses.
		LimitedRemoteAddresses bool
		// Connection size. High, Medium, Low. How many are
		// opened at any given situation.
	}
}

func (n *Network) Compute() {
	for i := 0; i < n.Duration; i++ {
		go func() {
			netSockStat, err := procfs.NewNetSockstat()
			if err != nil {
				fmt.Println("ERROR: getting information from /proc/sockstat")
			}

			n.NetSockstats = append(n.NetSockstats, netSockStat)
		}()

		// TODO(abhiyerra): All of these should also read IPv6 and UDP.

		time.Sleep(time.Second)
	}
}

func (n *Network) mean(title string, data []float64) {
	avg, err := stats.Mean(data)
	if err != nil {
		fmt.Println("Avg error")
		return
	}

	fmt.Println(title, avg)
}

func (n *Network) Avg() {
	var (
		usedData  []float64
		tcpTwData []float64
	)

	for _, i := range n.NetSockstats {
		usedData = append(usedData, float64(i.Sockets.Used))
		tcpTwData = append(tcpTwData, float64(i.TCP.Tw))
	}

	n.mean("Total:", usedData)
	n.mean("TCP TW:", tcpTwData)
}

func (n *Network) ParseArgs(args []string) {
	flags := flag.NewFlagSet(n.CmdName, flag.ContinueOnError)
	flags.IntVar(&n.Duration, "duration", 60, "Duration to monitor in seconds. Defaults to 60 seconds.")

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}
}

func (n *Network) Run() error {
	n.Compute()
	n.Avg()

	// TODO: Are the network connections to the same place?
	// TODO: Are

	return nil
}

func NewNetwork(cmdName string) *Network {
	return &Network{
		CmdName: cmdName,
	}
}

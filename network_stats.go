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
	"fmt"
	"time"

	"flag"
	"github.com/abhiyerra/procfs"
	"github.com/montanaflynn/stats"
	"os"
)

type NetworkStats struct {
	Duration     int
	NetSockstats []procfs.NetSockstat
}

func (n *NetworkStats) ComputeNetworkStats() {
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

func (n *NetworkStats) mean(title string, data []float64) {
	avg, err := stats.Mean(data)
	if err != nil {
		fmt.Println("Avg error")
		return
	}

	fmt.Println(title, avg)
}

func (n *NetworkStats) Avg() {
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

func (n *NetworkStats) ParseArgs(args []string) {
	flags := flag.NewFlagSet(CmdName, flag.ContinueOnError)
	flags.IntVar(&n.Duration, "duration", 60, "Duration to monitor in seconds. Defaults to 60 seconds.")

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}
}

func (n *NetworkStats) Run() error {
	n.ComputeNetworkStats()
	n.Avg()

	return nil
}

func NewNetworkStats() *NetworkStats {
	return &NetworkStats{}
}

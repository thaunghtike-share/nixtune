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
	"fmt"
	"time"

	"flag"
	"github.com/abhiyerra/procfs"
	"github.com/montanaflynn/stats"
	"os"
)

type Network struct {
	Duration     int
	NetSockstats []procfs.NetSockstat
}

func (n *Network) ComputeNetwork(duration int) {
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

func (n *Network) mean(data []float64) {
	avg, err := stats.Mean(data)
	if err != nil {
		fmt.Println("Avg error")
		return
	}

	fmt.Println("Avg Used:", avg)
}

func (n *Network) Avg() {
	var (
		usedData  []float64
		tcpTwData []float64
	)

	for _, i := range n.NetSockstats {
		usedData = append(usedData, float64(i.Sockets.Used))
		tcpTwData = append(tcpTwData, float64(i.TCP.InUse))
	}

	n.mean(usedData)
	n.mean(tcpTwData)
}

func (n *Network) ParseArgs(args []string) {
	flags := flag.NewFlagSet(CmdName, flag.ContinueOnError)
	flags.IntVar(&n.Duration, "duration", 60, "Duration to monitor in seconds.")

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}
}

func (n *Network) Run() error {
	n.Avg()

	return nil
}

func NewNetwork() *Network {
	return &Network{}
}

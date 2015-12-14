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
	"github.com/abhiyerra/procfs"
	"time"
)

type Network struct {
	Used     []int64
	InUse    []int64
	TimeWait []int64
}

func (n *Network) ComputeNetwork(duration int) {
	for i := 0; i < duration; i++ {
		netSockStat, err := procfs.NewNetSockstat()
		if err != nil {
			fmt.Println("Error getting information from /proc/sockstat")
		}

		// All of these should also read IPv6 and UDP.
		n.Used = append(n.Used, netSockStat.Sockets.Used)
		n.InUse = append(n.InUse, netSockStat.TCP.InUse)
		n.TimeWait = append(n.TimeWait, netSockStat.TCP.Tw)

		time.Sleep(time.Second)
	}
}

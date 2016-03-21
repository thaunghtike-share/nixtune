/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package network

import (
//	"github.com/acksin/procfs"
)

/*

TODO:
Network Access ALgorithm for Acksin STRUM

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

// Total: 804 (kernel 0)
// TCP:   410 (estab 327, closed 75, orphaned 0, synrecv 0, timewait 71/0), ports 0

// Transport Total     IP        IPv6
// *	  0         -         -
// RAW	  0         0         0
// UDP	  15        9         6
// TCP	  335       333       2
// INET	  350       342       8
// FRAG	  0         0         0

type TCPStat struct {
	Established int64
	Closed      int64
	Orphaned    int64
	Synrecv     int64
	Timewait    int64
}
type Network struct {
	Total int64

	// TODO Implement RAW sockets
	RAW struct{}
	// TODO Implement UDP sockets
	UDP struct{}

	TCP struct {
		Total int64
		V4    TCPStat
		V6    TCPStat
	}
}

func (n *Network) tcpIPv4() {
}

func (n *Network) tcpIPv6() {
}

func New() *Network {
	n := &Network{}

	n.tcpIPv4()
	n.tcpIPv6()

	return n
}

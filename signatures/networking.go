/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

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

type networkConfig struct {
}

func newNetworkConfig() *networkConfig {
	return &networkConfig{}
}

func (c *networkConfig) GetEnv() map[string]string {
	// No-op

	return nil
}

// Many of these settings were from the following places:
//   - http://vincent.bernat.im/en/blog/2014-tcp-time-wait-state-linux.html
//   - https://rtcamp.com/tutorials/linux/sysctl-conf/
//
// TODO: These setting are sort of set in stone but I feel that they
// can adapt as the system is being used. We don't have to set them to
// the values but we can migrate and change as we learn more about the
// system and tune it appropriately.

func (c *networkConfig) GetSysctl() map[string]string {
	sysctl := make(map[string]string)

	// tcp_fin_timeout - INTEGER
	//         Time to hold socket in state FIN-WAIT-2, if it was closed
	//         by our side. Peer can be broken and never close its side,
	//         or even died unexpectedly. Default value is 60sec.
	//         Usual value used in 2.2 was 180 seconds, you may restore
	//         it, but remember that if your machine is even underloaded WEB server,
	//         you risk to overflow memory with kilotons of dead sockets,
	//         FIN-WAIT-2 sockets are less dangerous than FIN-WAIT-1,
	//         because they eat maximum 1.5K of memory, but they tend
	//         to live longer. Cf. tcp_max_orphans.
	sysctl["net.ipv4.tcp_fin_timeout"] = "15"

	// On Linux, the client port has a range of about 30,000 ports. This
	// means that only 30,000 connections can be established between the
	// web server and the load-balancer every minute, so about 500
	// connections per second. We can increase the amount of available
	// ports.
	sysctl["net.ipv4.ip_local_port_range"] = "1024 65535"

	// 16MB per socket.
	sysctl["net.core.rmem_max"] = "16777216"
	sysctl["net.core.wmem_max"] = "16777216"

	// Increase the number syn requests allowed.
	sysctl["net.ipv4.tcp_max_syn_backlog"] = "20480"
	sysctl["net.ipv4.tcp_syncookies"] = "1"

	// The maximum number of "backlogged sockets".
	sysctl["net.core.somaxconn"] = "16096"

	sysctl["net.core.netdev_max_backlog"] = "30000"
	// Maximal number of timewait sockets held by the system
	// simultaneously. If this number is exceeded time-wait socket
	// is immediately destroyed and a warning is printed. This
	// limit exists only to prevent simple DoS attacks, you must
	// not lower the limit artificially, but rather increase it
	// (probably, after increasing installed memory), if network
	// conditions require more than the default value.
	sysctl["net.ipv4.tcp_max_tw_buckets"] = "400000"
	sysctl["net.ipv4.tcp_no_metrics_save"] = "1"
	sysctl["net.ipv4.tcp_rmem"] = "4096 87380 16777216"
	sysctl["net.ipv4.tcp_syn_retries"] = "2"
	sysctl["net.ipv4.tcp_synack_retries"] = "2"
	sysctl["net.ipv4.tcp_wmem"] = "4096 65536 16777216"

	// Amount of memory to keep free. Don't want to make this too high as
	// Linux will spend more time trying to reclaim memory.
	sysctl["vm.min_free_kbytes"] = "65536"

	// http://serverfault.com/questions/122679/how-do-ulimit-n-and-proc-sys-fs-file-max-differ
	sysctl["fs.file-max"] = "2097152"

	return sysctl
}

// http://cherokee-project.com/doc/other_os_tuning.html
// https://easyengine.io/tutorials/linux/sysctl-conf/

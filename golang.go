package main

func golangConfig(k *KnightAgent) (sc *SystemConfig) {
	sc = &SystemConfig{
		Sysctl: make(map[string]string),
		Env:    make(map[string]string),
		Files:  make(map[string]FileChange),
	}

	// Set the value of GOGC to be really high.

	// TODO: Consider how this is being used as part of a bigger
	// setting. Based on RAM etc.
	sc.Env["GOGC"] = "2000"

	golangNetworkSysctl(sc.Sysctl)

	// http://serverfault.com/questions/122679/how-do-ulimit-n-and-proc-sys-fs-file-max-differ
	sc.Sysctl["fs.file-max"] = "2097152"
	sc.Files["/etc/security/limits.d/00_anatma_knight_limits.conf"] = FileChange{
		Content: "* - nofile unlimited",
		Append:  true,
	}

	return
}

// Many of these settings were from the following places:
//   - http://vincent.bernat.im/en/blog/2014-tcp-time-wait-state-linux.html
//   - https://rtcamp.com/tutorials/linux/sysctl-conf/
//
// TODO: These setting are sort of set in stone but I feel that they
// can adapt as the system is being used. We don't have to set them to
// the values but we can migrate and change as we learn more about the
// system and tune it appropriately.
func golangNetworkSysctl(sysctl map[string]string) {
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
	sysctl["net.core.somaxconn"] = "4096"

	sysctl["net.core.netdev_max_backlog"] = "4096"
	sysctl["net.ipv4.tcp_max_tw_buckets"] = "400000"
	sysctl["net.ipv4.tcp_no_metrics_save"] = "1"
	sysctl["net.ipv4.tcp_rmem"] = "4096 87380 16777216"
	sysctl["net.ipv4.tcp_syn_retries"] = "2"
	sysctl["net.ipv4.tcp_synack_retries"] = "2"
	sysctl["net.ipv4.tcp_wmem"] = "4096 65536 16777216"

	// Amount of memory to keep free. Don't want to make this too high as
	// Linux will spend more time trying to reclaim memory.
	sysctl["vm.min_free_kbytes"] = "65536"
}

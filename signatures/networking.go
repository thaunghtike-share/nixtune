package signatures

type Networking struct{}

func (f *Networking) GetProfile() *Profile {
	p := &Profile{
		Name:          "networking",
		Subscription:  OpenSubscription,
		Description:   "Settings for high performance networking",
		Documentation: "TODO: These setting are sort of set in stone but I feel that they can adapt as the system is being used. We don't have to set them to the values but we can migrate and change as we learn more about the system and tune it appropriately.",
		References: []string{
			"http://vincent.bernat.im/en/blog/2014-tcp-time-wait-state-linux.html",
			"https://rtcamp.com/tutorials/linux/sysctl-conf/",
			"https://fasterdata.es.net/host-tuning/linux/",
			"http://cherokee-project.com/doc/other_os_tuning.html",
			"https://easyengine.io/tutorials/linux/sysctl-conf/",
			"https://access.redhat.com/sites/default/files/attachments/20150325_network_performance_tuning.pdf",
		},
		ProcFS: f.procfs(),
		SysFS:  f.sysfs(),
	}

	p.Vars = make(map[string]interface{})
	p.Vars["nfConntrackMax"] = 200000

	return p
}

func (f *Networking) procfs() map[string]ProfileKV {
	p := make(map[string]ProfileKV)

	p["net.ipv4.tcp_fin_timeout"] = ProfileKV{
		Value:       "15",
		Description: "Usually, the Linux kernel holds a TCP connection even after it is closed for around two minutes. This means that there may be a port exhaustion as the kernel waits to close the connections. By moving the fin_timeout to 15 seconds we drastically reduce the length of time the kernel is waiting for the socket to get any remaining packets.",
	}

	p["net.ipv4.ip_local_port_range"] = ProfileKV{
		Value:       "1024 65535",
		Description: "On a typical machine there are around 28000 ports available to be bound to. This number can get exhausted quickly if there are many connections. We will increase this.",
	}

	p["net.core.rmem_max"] = ProfileKV{
		Value:       "16777216",
		Description: "The size of the receive buffer for all the sockets. 16MB per socket.",
	}

	p["net.core.wmem_max"] = ProfileKV{
		Value:       "16777216",
		Description: "The size of the buffer for all the sockets. 16MB per socket.",
	}

	p["net.ipv4.tcp_rmem"] = ProfileKV{
		Value:       "4096 87380 16777216",
		Description: "(min, default, max): The sizes of the receive buffer for the IP protocol.",
	}

	p["net.ipv4.tcp_wmem"] = ProfileKV{
		Value:       "4096 65536 16777216",
		Description: "(min, default, max): The sizes of the write buffer for the IP protocol.",
	}

	p["net.ipv4.tcp_max_syn_backlog"] = ProfileKV{
		Value:       "20480",
		Description: "Increase the number syn requests allowed. Sets how many half-open connections to backlog queue",
	}

	p["net.ipv4.tcp_syncookies"] = ProfileKV{
		Value:       "1",
		Description: "Security to prevent DDoS attacks. http://cr.yp.to/syncookies.html",
	}

	p["net.ipv4.tcp_no_metrics_save"] = ProfileKV{
		Value:       "1",
		Description: "TCP saves various connection metrics in the route cache when the connection closes so that connections established in the near future can use these to set initial conditions. Usually, this increases overall performance, but may sometimes cause performance degradation.",
	}

	p["net.core.somaxconn"] = ProfileKV{
		Value:       "16096",
		Description: "The maximum number of queued sockets on a connection.",
	}

	p["net.core.netdev_max_backlog"] = ProfileKV{
		Value:       "30000",
		Description: "The number of incoming connections on the backlog queue. The maximum number of packets queued on the INPUT side.",
	}

	p["net.ipv4.tcp_max_tw_buckets"] = ProfileKV{
		Value:       "400000",
		Description: "Increase the tcp-time-wait buckets pool size to prevent simple DOS attacks",
	}

	p["net.ipv4.tcp_syn_retries"] = ProfileKV{
		Value:       "2",
		Description: "Number of times initial SYNs for a TCP connection attempt will be retransmitted for outgoing connections.",
	}

	p["net.ipv4.tcp_synack_retries"] = ProfileKV{
		Value:       "2",
		Description: "This setting determines the number of SYN+ACK packets sent before the kernel gives up on the connection",
	}

	p["net.netfilter.nf_conntrack_max"] = ProfileKV{
		Value:       "{{ index .Vars \"nfConntrackMax\" }}",
		Description: "The max is double the previous value. https://wiki.khnet.info/index.php/Conntrack_tuning",
	}

	p["net.ipv4.tcp_tw_reuse"] = ProfileKV{
		Value:       "1",
		Description: "",
	}

	return p
}

func (f *Networking) sysfs() map[string]ProfileKV {
	p := make(map[string]ProfileKV)
	p["/sys/module/nf_conntrack/parameters/hashsize"] = ProfileKV{
		Value: "{{ divide (index .Vars \"nfConntrackMax\") 4 }}",
	}

	return p
}

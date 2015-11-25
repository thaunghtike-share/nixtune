#!/usr/bin/env bash

# http://vincent.bernat.im/en/blog/2014-tcp-time-wait-state-linux.html

# echo "" >>

# tcp_fin_timeout - INTEGER
#         Time to hold socket in state FIN-WAIT-2, if it was closed
#         by our side. Peer can be broken and never close its side,
#         or even died unexpectedly. Default value is 60sec.
#         Usual value used in 2.2 was 180 seconds, you may restore
#         it, but remember that if your machine is even underloaded WEB server,
#         you risk to overflow memory with kilotons of dead sockets,
#         FIN-WAIT-2 sockets are less dangerous than FIN-WAIT-1,
#         because they eat maximum 1.5K of memory, but they tend
#         to live longer. Cf. tcp_max_orphans.

net.ipv4.tcp_fin_timeout=15


# On Linux, the client port is by default allocated in a port range of
# about 30,000 ports (this can be changed by tuning
# net.ipv4.ip_local_port_range). This means that only 30,000 connections
# can be established between the web server and the load-balancer every
# minute, so about 500 connections per second.
net.ipv4.ip_local_port_range = 1024 65535

#
# 16MB per socket - which sounds like a lot, but will virtually never
# consume that much.
#
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216

# Increase the number of outstanding syn requests allowed.
# c.f. The use of syncookies.
net.ipv4.tcp_max_syn_backlog='20480'
net.ipv4.tcp_syncookies = 1

# The maximum number of "backlogged sockets".  Default is 128.
net.core.somaxconn = 4096


net.core.netdev_max_backlog='4096'
net.ipv4.tcp_max_tw_buckets='400000'
net.ipv4.tcp_no_metrics_save='1'
net.ipv4.tcp_rmem='4096 87380 16777216'
net.ipv4.tcp_syn_retries='2'
net.ipv4.tcp_synack_retries='2'
net.ipv4.tcp_wmem='4096 65536 16777216'
vm.min_free_kbytes='65536'

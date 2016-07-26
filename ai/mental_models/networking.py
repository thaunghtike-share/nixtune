# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

class Networking(object):
    """
    References:

    http://vincent.bernat.im/en/blog/2014-tcp-time-wait-state-linux.html
    https://rtcamp.com/tutorials/linux/sysctl-conf/
    https://fasterdata.es.net/host-tuning/linux/
    http://cherokee-project.com/doc/other_os_tuning.html
    https://easyengine.io/tutorials/linux/sysctl-conf/
    https://access.redhat.com/sites/default/files/attachments/20150325_network_performance_tuning.pdf
    """

    def __init__(self, strum):
        self.strum = strum

        self.vars = {
            'nfConntrackMax': 200000
        }

    def procfs_net_ipv4_tcp_fin_timeout(self):
        """
        Usually, the Linux kernel holds a TCP connection even after it
        is closed for around two minutes. This means that there may be
        a port exhaustion as the kernel waits to close the
        connections. By moving the fin_timeout to 15 seconds we
        drastically reduce the length of time the kernel is waiting
        for the socket to get any remaining packets.
        """

        return {
            "/proc/sys/net/ipv4/tcp_fin_timeout": "15"
        }

    def procfs_net_ipv4_ip_local_port_range(self):
        """
        On a typical machine there are around 28000 ports available to be
        bound to. This number can get exhausted quickly if there are many
        connections. We will increase this.
        """

        return {
            "net.ipv4.ip_local_port_range": "1024 65535",
        }

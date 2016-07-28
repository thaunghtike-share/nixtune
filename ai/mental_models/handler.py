# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import json
import psycopg2
import sys

from memory import Memory
from networking import Networking

class Strum(object):
    def __init__(self, id):
        self.ID = id

        self.config = json.load(open('config.json', 'r'))
        self.conn = psycopg2.connect(self.config['database'])

        cur = self.conn.cursor()

        cur.execute("SELECT data FROM strum_stats where id = %s", (id,))

        self.stats = cur.fetchone()[0]

        self.conn.commit()

        cur.close()

    def close(self):
        self.conn.close()

def handler(event, context):
    """
    Run on AWS Lambda.
    """

    strum = Strum(event['ID'])

    memory = Memory(strum)
    memory.is_swapping()
    memory.is_under_utilized()
    memory.procfs_vm_swappiness()
    memory.procfs_proc_min_free_kbytes()
    memory.sysfs_mm_transparent_hugepages()

    networking = Networking(strum)
    networking.procfs_net_ipv4_tcp_fin_timeout()
    networking.procfs_net_ipv4_ip_local_port_range()
    networking.procfs_net_core_rmem_max()
    networking.procfs_net_core_wmem_max()
    networking.procfs_net_ipv4_tcp_rmem()
    networking.procfs_net_ipv4_tcp_wmem()
    networking.procfs_net_ipv4_tcp_max_syn_backlog()
    networking.procfs_net_ipv4_tcp_syncookies()
    networking.procfs_net_ipv4_tcp_no_metrics_save()
    networking.procfs_net_core_somaxconn()
    networking.procfs_net_core_netdev_max_backlog()
    networking.procfs_net_ipv4_tcp_max_tw_buckets()
    networking.procfs_net_ipv4_tcp_syn_retries()
    networking.procfs_net_ipv4_tcp_synack_retries()
    networking.procfs_net_netfilter_nf_conntrack_max()
    networking.procfs_net_ipv4_tcp_tw_reuse()
    networking.sysfs_nf_conntrack_hashsize()

    strum.close()

    return {
        'Message': "OK"
    }

if __name__ == "__main__":
    print handler({
        'ID': sys.argv[1]
    }, None)

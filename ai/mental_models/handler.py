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
    The handler works to so that it can also run on AWS Lambda.
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

    strum.close()

    return {
        'Message': "OK"
    }

if __name__ == "__main__":
    print handler({
        'ID': sys.argv[1]
    }, None)

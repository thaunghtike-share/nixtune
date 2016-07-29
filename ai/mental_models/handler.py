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

        cur.close()

    def close(self):
        self.conn.close()

def handler(event, context):
    """
    Run on AWS Lambda.
    """

    strum = Strum(event['ID'])

    memory = Memory(strum)
    memory.ai_features()
    memory.procfs_features()
    memory.sysfs_features()

    networking = Networking(strum)
    networking.ai_features()
    networking.procfs_features()
    networking.sysfs_features()

    strum.close()

    return {
        'Message': "OK"
    }

if __name__ == "__main__":
    print handler({
        'ID': sys.argv[1]
    }, None)

# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import sys; sys.path.append("site-packages")
import json
import pg8000
import boto3

from memory import Memory
from networking import Networking

class Autotune(object):
    def __init__(self, config_file):
        self.config = json.load(open(config_file, 'r'))
        self.conn = pg8000.connect(
            user=self.config['username'],
            password=self.config['password'],
            database=self.config['database'],
            host=self.config['host'],
            ssl=self.config['ssl']
        )

    def get_ids(self):
        """
        Get the entire list of stats ids from the database.
        """

        cur = self.conn.cursor()
        cur.execute("SELECT id::varchar FROM autotune_stats")
        ids = [i[0] for i in cur.fetchall()]
        cur.close()

        return ids

    def get_id(self, id):
        self.ID = id

        cur = self.conn.cursor()
        cur.execute("SELECT data FROM autotune_stats where id = %s", (id,))
        self.stats = cur.fetchone()[0]
        cur.close()

    def write_ai_features(self, features):
        cur = self.conn.cursor()
        cur.execute("UPDATE autotune_stats SET ai_features = %s WHERE id = %s", (json.dumps(features), self.ID))
        cur.close()
        self.conn.commit()

    def write_procfs_features(self, features):
        cur = self.conn.cursor()
        cur.execute("UPDATE autotune_stats SET procfs_features = %s WHERE id = %s", (json.dumps(features), self.ID))
        cur.close()
        self.conn.commit()

    def write_sysfs_features(self, features):
        cur = self.conn.cursor()
        cur.execute("UPDATE autotune_stats SET sysfs_features = %s WHERE id = %s", (json.dumps(features), self.ID))
        cur.close()
        self.conn.commit()

    def close(self):
        self.conn.close()

def run_model(config_file, event_id):
    print "Running {}".format(event_id)

    autotune = Autotune(config_file)
    autotune.get_id(event_id)

    memory = Memory(autotune)
    networking = Networking(autotune)

    ai_features = dict(memory.ai_features().items() + networking.ai_features().items())
    procfs_features = dict(memory.procfs_features().items() + networking.procfs_features().items())
    sysfs_features = dict(memory.sysfs_features().items() + networking.sysfs_features().items())

    autotune.write_ai_features(ai_features)
    autotune.write_procfs_features(procfs_features)
    autotune.write_sysfs_features(sysfs_features)

    autotune.close()

def run_upgrade(config_file, function_name):
    autotune = Autotune(config_file)
    client = boto3.client('lambda')

    for autotune_id in autotune.get_ids():
        client.invoke(
            FunctionName=function_name,
            InvocationType='Event',
            Payload=bytearray(json.dumps({'ID': autotune_id}), 'utf-8'),
        )

def handler(event, context):
    """
    Run on AWS Lambda.
    """

    config_file = "config.dev.json"
    if context is not None and context.function_name == "autotune-prod-mentalmodels":
        config_file = "config.prod.json"

    if event.has_key('ID'):
        run_model(config_file, event['ID'])
    elif event.has_key('Upgrade'):
        run_upgrade(config_file, context.function_name)

    return {
        'Message': "OK"
    }

if __name__ == "__main__":
    if sys.argv[1] == 'Upgrade':
        print handler({
            'Upgrade': True,
        }, None)
    else:
        print handler({
            'ID': sys.argv[1]
        }, None)

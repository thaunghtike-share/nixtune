# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import json
import os
import boto3

class Autotune(object):
    def __init__(self, config_file):
        self.config = json.load(open(config_file, 'r'))

        self.s3 = boto3.client('s3')

    def write_object(self, key, value):
        self.s3.put_object(
            Bucket=self.config['s3_bucket'],
            Key=os.path.join(self.ID, key),
            Body=value.encode('utf-8')
        )

    def get_id(self, id):
        self.ID = id

        resp = self.s3.get_object(
            Bucket=self.config['s3_bucket'],
            Key=os.path.join(self.ID, 'stats.json')
        )

        self.stats = json.loads(resp['Body'].read())

    def write_ai_features(self, features):
        self.write_object("ai_features.json", json.dumps(features))

    def write_procfs_features(self, features):
        self.write_object("procfs.json", json.dumps(features))

    def write_sysfs_features(self, features):
        self.write_object("sysfs.json", json.dumps(features))

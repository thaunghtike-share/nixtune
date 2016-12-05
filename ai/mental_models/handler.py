# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import sys; sys.path.append("site-packages")
import os

import machine

def handler(event, context):
    """
    Run on AWS Lambda.
    """

    config_file = "config.dev.json"
    if context is not None and context.function_name == "acksin-prod-mentalmodels":
        config_file = "config.prod.json"

    machine.run_model(config_file, event['ID'])

    return {
        'Message': "OK"
    }

if __name__ == "__main__":
    print handler({
        'Machine': True,
        'ID': sys.argv[1]
    }, None)

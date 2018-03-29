# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

from decorators import *
import mental_model

class IO(mental_models.MentalModel):
    """
    References:
      - http://www.brendangregg.com/linuxperf.html
    """

    def __init__(self, autotune):
        self.autotune = autotune

    @sysfs_feature
    def sysfs_block_queue_rq_afinity(self):
        return {
            "/sys/block/*/queue/rq_afinity": "2"
        }

    @sysfs_feature
    def sysfs_block_queue_scheduler(self):
        return {
            "/sys/block/*/queue/scheduler": "noop"
        }

    @sysfs_feature
    def sysfs_block_queue_read_ahead_kb(self):
        return {
            "/sys/block/*/queue/read_ahead_kb": "256",
        }

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

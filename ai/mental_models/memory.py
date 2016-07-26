# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

class Memory(object):
    def __init__(self, strum):
        self.strum = strum
        self.memory = self.strum.stats['System']['Memory']

    def is_swapping(self):
        """
        If The machine is swapping memory consider moving to a larger machine with more memory.
        """

        return self.memory['Physical']['Free'] == 0 and self.memory['Swap']['Used'] > 0

    def is_under_utilizied(self):
        """
        Linux uses some of the free memory for storing file buffers in
        memory. Let's see how much it caches and recommend an instance
        size.

        http://askubuntu.com/questions/198549/what-is-cached-in-the-top-command

        """
        if self.memory['Physical']['Free'] > 0:
            percent_used = self.memory['Physical']['Free'] / self.memory['Physical']['Total']
            return percent_used < 0.5

        return False

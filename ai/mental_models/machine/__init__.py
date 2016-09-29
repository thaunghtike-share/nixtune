# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

from autotune import Autotune

from memory import Memory
from networking import Networking

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

# Copyright 2016 Acksin, LLC. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from __future__ import absolute_import
from __future__ import division
from __future__ import print_function

import json
import csv
import tensorflow as tf
import numpy as np
import pandas as pd
import random

file = "config.json"
columns = ['Process',
def parse_json():
    # Here we will parse a JSON file with various Linux information
    # such as CPU, Networking, IO, Memory, Processes, Limits, Disks
    # and more

    with open(file) as f:
        data = json.load(f)

    train_file = csv.writer(open("train.csv", "wb+"))
    test_file = csv.writer(open("test.csv", "wb+"))

    # Let's look at and count the processes
    processes = data['Processes']
    index = 0

    for process in processes:
        print(process['Exe'])
        print(process['PID'])
        print(process['Memory']['Swap']['Size'])
        index += 1

    print("There are %d processes to account for."  % index)

    # Now let's look at the memory
    system = data['System']
    memory = system['Memory']

    # Physical
    phys_mem = memory['Physical']
    phys_free = phys_mem['Free']
    phys_used = phys_mem['Used']
    phys_total = phys_mem['Total']
    phys_cached = phys_mem['Cached']
    system_stats("Physical", phys_free, phys_used, phys_total, phys_cached)

    # Swap
    swap_mem = memory['Swap']
    swap_free = swap_mem['Free']
    swap_used = swap_mem['Used']
    swap_total = swap_mem['Total']
    swap_cached = swap_mem['Cached']
    system_stats("Swap", swap_free, swap_used, swap_total, swap_cached)

    # Virtual
    virt_mem = memory['Virtual']
    #virt_free = virt_mem['Free']
    virt_used = virt_mem['Used']
    virt_total = virt_mem['Total']
    virt_chunk = virt_mem['Chunk']
    #system_stats("Virtual", virt_free, virt_used, virt_total, virt_cached)

    disk_devices = system['Disk']['BlockDevices']

    for device in disk_devices:
        print(device['name'])
        if 'children' not in device:
            print("Has no children")
        else:
            print("Has Children")

def system_stats(mem_type, free, used, total, cached):
    print("Memory type: " + mem_type)
    print("Percent free is: %d" % ((free/total) * 100) + "%")
    print("Percent used is: %d" % ((used/total) * 100) + "%")
    print("Percent cached is: %d" % ((cached/total) * 100) + "%")
    print(type(free))
    print(type(used))
    print(type(total))

parse_json()

# Acksin STRUM

Tool to gather quick stats about a machine including metrics on CPU,
IO, Memory, Networking, etc. to help quickly debug and locate common
problems.

## Use Case

DevOps people turn to tools ls =vmstat=, =free=, =/proc= fs, etc. to
quickly figure out what the issues are with a Linux and UNIX
machines. However, as common as these tools are they are still limited
insofar as they need to be strung togehter to give a complete picture
of what is happening.

If a MySQL machine is high CPU your workflow may be running =top= to
see what the CPUs are like, =free= to see if the machine is swapping,
=ss= to see if there are a lot of connections, etc.

Strum will give you a birds eye view of the system quickly.

## Usage

```
sudo strum
```

If you just want infor on a single PID run the following:

```
sudo strum [pid]
```

## Completed

 - [X] Memory
 - [ ] Networking
 - [ ] IO
 - [ ] CPU
 - [ ] FDs
 - [ ] Process
   - [ ] Memory
   - [ ] Networking
   - [ ] IO
   - [ ] CPU
   - [ ] FDs



## Authors

 Abhi Yerra @abhiyerra

## License

Copyright (C) 2015 Acksin <hey@acksin.com>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.

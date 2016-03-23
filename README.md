[![Build Status](https://travis-ci.org/acksin/strum.svg?branch=master)](https://travis-ci.org/acksin/strum)

# Acksin STRUM

Tool to diagnose issues on the machine quickly. It includes system and
process information about CPU, IO, Memory, Networking, Limits, etc. to
help debug and assess problems.

## Use Case

System Admins turn to tools like `vmstat`, `free`, `top`, `ps`,
etc. to quickly figure out what the issues are with a Linux and UNIX
machines. However, as common as these tools are they are still limited
insofar as they need to be strung togehter to give a complete picture
of what is happening.

If a MySQL machine is high CPU your workflow may be running `top` to
see what the CPUs are like, `free` to see if the machine is swapping,
`ss` to see if there are a lot of connections, etc.

STRUM will give you a birds eye view of the system and a process
quickly.

## Usage

While the tool can be used without sudo there may be information that
is left out because of lack of access. As such it is better to run the
tool as `sudo`.

```
sudo strum
```

If you just want info on a single PID run the following:

```
sudo strum [pid]
```

## Authors

 Abhi Yerra @abhiyerra

## License

Copyright (C) 2015 Acksin <hey@acksin.com>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.

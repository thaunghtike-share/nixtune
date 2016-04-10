![](website/autotune.png)

# Acksin Autotune

Autotune the Linux kernel to get optimal server performance.

Acksin Autotune gives you the Linux kernel turning and environment
variables to get maximum performance for you applications and to
utilize the machine to its fullest.

Autotune figures out the best settings for you to use when running
your applications.

## Usage

Using the signatures is straightforward. Just pass the names of the
signatures as part of the command and it will give you the appropriate
values to tune.

### Output the Raw signature as JSON

```
autotune sig golang
```

### ProcFS changes

```
autotune sig --procfs golang
```

### SysFS changes

```
autotune sig --sysfs golang
```

### Environment Variables

```
autotune sig --env golang
```

## Documentation

Documentation as well as the descriptions of each of the changes for
the signatures that are taken into account are located
[https://www.acksin.com/autotune/docs](Acksin Autotune Docs) website

## Open Signatures

 - apache
 - golang
 - haproxy
 - nginx
 - nodejs
 - postgresql

## Pro Signatures

 - [ ] docker
 - [ ] java
 - [ ] memcached
 - [ ] mod_passenger
 - [ ] mysql
 - [ ] php
 - [ ] python
 - [ ] redis
 - [ ] ruby
 - [ ] rails

## Premium Signatures
 - [ ] cassandra
 - [ ] hadoop
 - [ ] mesos
 - [ ] mongodb
 - [ ] mongodb-wiredtiger
 - [ ] spark


## License

Copyright (C) 2016 Acksin <hey@acksin.com>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.

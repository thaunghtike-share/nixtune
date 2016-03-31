![](website/autotune.png)

# Acksin Autotune

Autotune the Linux kernel to get optimal server performance.

Acksin Autotune tunes the Linux kernel and other environment variables
to achieve maximum performance for you applications. The goal is to
support most common use cases such a Golang, Node, Java, Various
databases, and so on. These are an alternate to manually tuning the
kernel for performance which is a matter of trial and error.

The tool will attempt to figure out the best settings and if there is
a degration will rollback changes to their existing settings.

## Usage

Using the signatures is straightforward. Just pass the names of the
signatures as part of the command and it will tune it correctly.

```
autotune signature golang
```

### Signatures

 - [X] golang
 - [X] nodejs
 - [X] nginx
 - [X] haproxy
 - [X] apache
 - [X] postgresql

### Future Signatures

 - [ ] mysql
 - [ ] redis
 - [ ] mongodb-wiredtiger
 - [ ] mongodb
 - [ ] cassandra
 - [ ] java
 - [ ] hadoop
 - [ ] docker
 - [ ] mesos
 - [ ] spark

## Authors

 Abhi Yerra @abhiyerra

## License

Copyright (C) 2016 Acksin <hey@acksin.com>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.

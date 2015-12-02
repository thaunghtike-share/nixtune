# Anatma Knight

Auto tune the Linux kernel for your application.

Anatma Knight is a tool that auto tunes the Linux kernel and other
environment variables to achieve maximum performance for you
applications. The goal is to support most common use cases such a
Golang, Node, Java, Various databases, and so on. These are an
alternate to manually tuning the kernel for performance which is a
matter of trial and error.

The tool will attempt to figure out the best settings and if there is
a degration will rollback changes to their existing settings.

## Usage

```
knight -signature=golang
```

### Available Signatures

 - golang
 - nodejs
 - nginx
 - apache
 - postgresql
 - java

## Commercial License

I sell Anatma Knight Pro which provides a commercial-friendly license
as well as additional support. This allows me to produce high quality
open source code and be able to support myself in its
development. Please see the Anatma Knight homepage for more details.

## License

Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

## Author

 Abhi Yerra @abhiyerra

#

<a href="https://travis-ci.org/opszero/opszero"><img src="https://travis-ci.org/opszero/opszero.svg?branch=master" /></a>
<a href="https://godoc.org/github.com/opszero/opszero"><img src="https://godoc.org/github.com/opszero/opszero?status.svg" alt="GoDoc"></a>

## Introduction

opszero is Automatic Transmission for Linux. It attempts to performance
enhance your Linux server based on the type of services that you run
on it making sure that Linux is pulling its weight.

## [Quick Start and Download](https://www.opszero.com/#quickstart)

opszero outputs its data in JSON to the command line. Run the
following command:

    sudo opszero output

opszero primarily runs as a daemon which regularily pushes diagnostics
to a central server. opszero runs a service called
[opszero Console](https://www.opszero.com/console/login?redirectTo=https://www.opszero.com/console/)
providing this capability. You can get the configuration on the opszero
Console or you can check out
[config.json.template](config.json.template) for agent
configuration. We will open source the server side in the near future.

Run the following:

    sudo opszero agent config.json

## Getting Started & Documentation

All documentation is on the [opszero website](https://www.opszero.com/).

## Developing opszero

opszero's command line portion is primarily written in Go whereas the
Machine Learning is written in Python. The code is split into a
couple different sections:

 - [Command Line Tool](stats): Collects stats from the System and
   Containers.
 - [Mental Models](ai/mental_models): Take System stats and creates
   models for AI. Currently this is a program that runs on AWS Lambda.

### Primary Dependencies

The primary dependency of opszero is the
[ProcFS Library](https://github.com/opszero/procfs) we use.  Any code
that needs to read from ProcFS should go there. Most of the Command
Line App is a wrapper for that library. In the future we will have
similar dependencies for SysFS. In addition to that we use the Go
libraries provided by the Cloud providers.

### Deploying opszero

opszero has several components that need to be deployed to make a
complete system. This includes the Command Line Tool, Mental Models,
AI, Console, and Server.

#### Command Line Tool

To build the command line tool run the following:

```
make deps
make build
```

## License

Copyright (C) 2016 opszero <hey@opszero.com>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at <http://mozilla.org/MPL/2.0/>.
